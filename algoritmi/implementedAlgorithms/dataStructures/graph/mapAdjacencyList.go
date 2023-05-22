package graph

import (
	"fmt"
)

/*
	Questa implementazione è realizzata tramite una mappa.
	Questa implementazione rappresenta dei grafi ORIENTATI.
	Struct che rappresenta un grafo formato da una mappa che associa un intero a una lista d'interi, in questo modo per ogni
	elemento del grafo corrisponderà una list (di adiacenza) contenente tutti i nodi collegati all'intero in questione.
	La lista di adiacenza ha un costo di θ(n+m) in quanto il numero di liste di nodi adiacenti (m) dipende dal numero di nodi
	presenti nel grafo (n).
	In questo caso il numero di slice d'interi dipendono dal numero di elementi presenti nella mappa.
	Un problema della lista di adiacenza è il fatto che risulta costoso calcolare gli archi entranti in un nodo (per grafi orientati).
*/

type AdjMapListGraph struct {
	verts int
	data  map[int][]int
}

/*
	Genera un grafo contenente 'verts' nodi vuoti
*/

func NewMapAdjListGraph(verts int) *AdjMapListGraph {
	graph := new(AdjMapListGraph)
	graph.data = make(map[int][]int, verts)
	graph.verts = verts
	fmt.Println(graph)
	return graph
}

/*
	Permette di riempire il grafo con i dati passati
*/

func fillMapAdjListGraph(graph *AdjMapListGraph, node int, adjacent int) int {
	currVerst := len(graph.data) // numero corrente di nodi presenti nel grafo
	insertedNode := 0            // nodi inseriti in questa chiamata
	if node == adjacent {        // controllo per far si che un nodo non possa essere vicino di se stesso
		return 0
	}
	if currVerst < graph.verts { // se il numero di nodi presenti nel grafo è minore del numero massimo procedi
		if graph.data[node] == nil { // se il nodo non esiste crealo e metterlo nel grafo
			graph.data[node] = make([]int, 0)
			insertedNode++ // incremento numero di nodi aggiunti in questa chiamata
		}
		if graph.data[adjacent] == nil { // se il nodo vicino di 'node' non è presente nel grafo aggiungilo
			graph.data[adjacent] = make([]int, 0)
			insertedNode++ // incremento numero di nodi aggiunti in questa chiamata
		}
		// aggiungi il nodo 'adjacent' alla lista di adiacenza di 'node'
		graph.data[node] = append(graph.data[node], adjacent)
	} else { // altrimenti
		if graph.data[node] != nil && graph.data[adjacent] != nil { // verifica che il nodo 'node' e 'adjacent' esistano già
			// se si aggiungi 'adjacent' alla lista di adiacenza di 'node', in questo caso verranno connessi solo nodi già
			// presenti nel grafo, e non verranno aggiunti nodi nuovi
			graph.data[node] = append(graph.data[node], adjacent)
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

func ReadMapAdjListGraph(graph *AdjMapListGraph) {
	fmt.Println("Inserisci i nodi del grafo con i suoi nodi adiacenti nel seguente formato (per interrompere basta inserire una qualsiasi lettera): ")
	fmt.Println("nodo:adiacente")

	var node, adjacent int

	for i := 0; i <= graph.verts; {
		_, err := fmt.Scanf("%d:%d", &node, &adjacent)
		if err != nil {
			break
		}
		i = i + fillMapAdjListGraph(graph, node, adjacent)
	}
}

/*
	Permette di stampare il grafo nel formato
	nodo1 => vicino1-1, vicino1-2 ...
	nodo2 => vicino2-1, vicino2-2 ...
	nodo3 => vicino3-1, vicino3-2 ...
	...
*/

func PrintMapAdjListGraph(graph *AdjMapListGraph) {
	for node := range graph.data {
		fmt.Printf("%d => ", node)
		for _, adj := range graph.data[node] {
			fmt.Printf("%d ", adj)
		}
		fmt.Println()
	}
}
