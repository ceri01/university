package graph

import (
	util "algoritmi/algoritmi/implementedAlgorithms/dataStructures"
	"algoritmi/algoritmi/implementedAlgorithms/dataStructures/linkedList"
	"algoritmi/algoritmi/implementedAlgorithms/dataStructures/tree"
	"fmt"
)

/*
	Questa implementazione è realizzata tramite una lista (in go è possibile anche usare delle slice in modo semplice).
	Per avere tutti i tipi di grafo questa implementazione rappresenta dei grafi NON ORIENTATI.
	Inoltre in questo caso le chiavi dei nodi devono andare da zero a n dove n è il numero massimo di nodi che il grafo può
	contenere, questo perchè è un limite dell'implementazione, bisognerebbe passare a un'implementazione più complessa oppure
	utilizzare una mappa.
	Struct che rappresenta un grafo formato da una lista che associa un intero (la posizione) a una lista d'interi, in questo modo per ogni
	elemento del grafo corrisponderà una lista (di adiacenza) contenente tutti i nodi collegati all'intero in questione.
	Inoltre sono presenti due interi che indicano rispettivamente il numero di nodi che il grafo possiede e il numero di nodi
	già popolati.
	La lista di adiacenza ha un costo in termini di spazio di θ(n+m) in quanto il numero di liste di nodi adiacenti (m)
	dipende dal numero di nodi presenti nel grafo (n).
	In questo caso il numero di liste d'interi dipende dalla lunghezza della slice.
	Un problema della lista di adiacenza è il fatto che risulta costoso calcolare gli archi entranti in un nodo (per grafi orientati).
*/

type AdjListGraph struct {
	verts int
	// Ogni posizione i rappresenta un nodo (intero) e ogni lista associata a una posizione contiene l'insieme dei nodi
	// collegati a esso tramite un arco uscente i (se grafo orientato)
	data          []*linkedList.DoubleLinkedList
	vertOccupated int // Permette di tenere traccia di quanti vertici sono stati popolati
}

/*
	Genera un grafo contenente 'verts' nodi vuoti
*/

func NewAdjListGraph(verts int) *AdjListGraph {
	graph := new(AdjListGraph)
	graph.data = make([]*linkedList.DoubleLinkedList, verts)
	graph.verts = verts
	graph.vertOccupated = 0
	return graph
}

/*
	Permette di riempire il grafo con i dati passati, ritorna il numero di nodi adiacenti inseriti.
	Essendo l'implementazione di un arco non orientato quando viene inserito l'arco (u, v) verrà inserito anche l'arco (v, u)
	Il costo di questa operazione è lineare perchè si verifica se il nodo è già stato inserito, in quel caso non verrà inserito
	il valore due volte nella lista di adiacenza (per cercare si deve scorrere tutta la lista di adiacenza usando una funzione
	con costo O(n)).

*/

func fillAdjListGraph(graph *AdjListGraph, node int, adjacent int) int {
	insertedNode := 0     // nodi inseriti in questa chiamata
	if node == adjacent { // controllo per far si che un nodo non possa essere vicino di se stesso
		return 0
	} else if node >= graph.verts || adjacent >= graph.verts || node < 0 || adjacent < 0 {
		fmt.Println("Errore di inserimento, i nodi devono avere un valore minore di ", graph.verts)
		return 0
	}
	if graph.vertOccupated < graph.verts { // se il numero di nodi presenti nel grafo è minore del numero massimo procedi
		if graph.data[node] == nil { // se il nodo non esiste crealo e metterlo nel grafo
			graph.data[node] = new(linkedList.DoubleLinkedList)
			insertedNode++ // incremento numero di nodi aggiunti in questa chiamata
		}
		if graph.data[adjacent] == nil { // se il nodo vicino di 'node' non è presente nel grafo aggiungilo
			graph.data[adjacent] = new(linkedList.DoubleLinkedList)
			insertedNode++ // incremento numero di nodi aggiunti in questa chiamata
		}
		// aggiungi il nodo 'adjacent' alla lista di adiacenza di 'node' e viceversa (se non presenti)
		isPresent, _ := graph.data[node].SearchByKey(adjacent)
		if !isPresent {
			graph.data[node].InsertTail(adjacent)
			graph.data[adjacent].InsertTail(node)
		}
	} else { // altrimenti
		if graph.data[node] != nil && graph.data[adjacent] != nil { // verifica che il nodo 'node' e 'adjacent' esistano già
			// se si aggiungi 'adjacent' alla lista di adiacenza di 'node', in questo caso verranno connessi solo nodi già
			// presenti nel grafo, e non verranno aggiunti nodi nuovi, ma se questa connessione esiste già non verrà ripetuta.
			isPresent, _ := graph.data[node].SearchByKey(adjacent)
			if !isPresent {
				graph.data[node].InsertTail(adjacent)
				graph.data[adjacent].InsertTail(node)
			}
		} else { // altrimenti se si cerca di aggiungere un nuovo nodo verrà specificato che il limite è stato raggiunto
			fmt.Println("Numero massimo di nodi raggiunto")
			insertedNode++
		}
	}
	return insertedNode
}

/*
	Permette di popolare il grafo tramite standard input.
	È necessario rispettare un formato per poter inserire i dati, se ciò non avviene non è garantito che la funzione
	operi come dovrebbe

	FORMATO:
	nodo:vicino
*/

func ReadAdjListGraph(graph *AdjListGraph) {
	fmt.Println("Inserisci i nodi del grafo con i suoi nodi adiacenti nel seguente formato (per interrompere basta inserire una qualsiasi lettera): ")
	fmt.Println("nodo:adiacente")

	var node, adjacent int

	for i := 0; i <= graph.verts; {
		_, err := fmt.Scanf("%d:%d", &node, &adjacent)
		if err != nil {
			break
		}
		i = i + fillAdjListGraph(graph, node, adjacent)
	}
}

/*
	Permette di stampare il grafo nel formato
	nodo1 => vicino1-1, vicino1-2 ...
	nodo2 => vicino2-1, vicino2-2 ...
	nodo3 => vicino3-1, vicino3-2 ...
	...
*/

func PrintAdjListGraph(graph *AdjListGraph) {
	for key, node := range graph.data {
		if node != nil {
			fmt.Printf("%d => ", key)
			_, val := node.SearchByPosition(0)
			for val != nil {
				fmt.Printf("%d ", val.Key)
				val = val.Next
			}
			fmt.Println()
		}
	}
}

/*
	Permette di eseguire una vista in profondità del grafo, restituisce un albero ricoprente del grafo costruito selezionando
	gli archi secondo l'ordine della visita.
*/

func DFS(graph *AdjListGraph, node int) *util.BiSearchTree {
	t := tree.CreateBinarySearchTree(1)
	visited := make(map[int]bool, 1)
	visited[node] = true
	fmt.Printf("%d ", node)
	dfsRec(graph, node, t, visited)
	fmt.Println()
	return t
}

/*
	Procedura ricorsiva usata per fare la visita in profondità
*/

func dfsRec(graph *AdjListGraph, node int, searchTree *util.BiSearchTree, visited map[int]bool) {
	visited[node] = true
	_, el := graph.data[node].SearchByPosition(0)
	for el != nil {
		if visited[el.Key] != true {
			fmt.Printf("%d ", el.Key)
			tree.InsertIterative(searchTree, el.Key)
			dfsRec(graph, el.Key, searchTree, visited)
		}
		el = el.Next
	}
}
