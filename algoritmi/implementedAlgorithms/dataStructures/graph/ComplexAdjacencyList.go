package graph

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
	Questo tipo d'implementazione cosi articolata è necessario dal momento in cui si ha a che fare con dei dati piuttosto
	complessi come il tipo vertex qui sotto, che comprende una chiave (stringa) e una serie di stringhe come dati. Tutte le altre
	implementazioni riguarderanno casi più semplici, ovvero grafi su interi, si sappia però che in base al tipo di dati da
	rappresentare l'implementazione può variare.
	Un esempio che può essere una via di mezzo tra questa implementazione e le altre è quello in cui si deve rappresentare
	i dati come semplici stringhe, in quel caso basterà creare un'implementazione in cui si sfrutterà una mappa dove la
	chiave è una stringa e il valore associato è una slice di stringhe che rappresenta la lista di adiacenza.
*/

/*
	Struct che rappresenta un vertice del grafo formato da una chiave (stringa) e da una slice di stringhe che rappresentano
	i dati associati al nodo.
*/

type Vertex struct {
	Data [3]string
	Key  string
}

/*
	Struct che rappresenta un grafo formato da una mappa che associa una stringa ad un vertice (per identificare i nodi del grafo).
	Inoltre è presente una mappa che associa a ogni vertice una slice di riferimenti a vertici che rappresenta gli archi tra un nodo e
	i suoi adiacenti.
*/

type Graph struct {
	nodes int
	vert  map[string]*Vertex
	Adj   map[Vertex][]*Vertex
}

/*
	Genera un grafo contenente 'verts' nodi vuoti
*/

func NewGraph(verts int) *Graph {
	graph := new(Graph)
	graph.vert = make(map[string]*Vertex, verts)
	graph.Adj = make(map[Vertex][]*Vertex, 1)
	graph.nodes = verts
	return graph
}

/*
	Popola il grafico con i nodi rappresentati dalla chiave 'node' e se presenti setta i dati dei nodi, inoltre riempie la
	lista di adiacenza tramite il parametro adjacent.
	Restituisce il numero di nodi che sono stati inseriti

	Considerando che in questa funzione non si itera su alcuna lista o mappa si potrebbe dire che la complessità di questa funzione
	è costante, ma siccome sono presenti delle append che inseriscono degli elementi nelle liste di adiacenza non è cosi banale valutare la complessità.
	Infatti il costo ammortizzato della append è di O(1), ma nel caso peggiore è O(n), e ciò dipende dal tipo di input, che in questo caso non sono dati
	primitivi ma stringhe. Quindi si può concludere che il costo di questa funzione è:

	Caso migliore: O(1)
	Caso peggiore O(n)

	Per quanto riguarda lo spazio anche qui dipende dall'input, in base alla lunghezza delle stringe in input avremo un
	consumo di spazio diverso (più le stringhe sono grandi maggiore è lo spazio occupato)
*/

func fillGraph(graph *Graph, node string, data [3]string, adjacent string) int {
	currVerst := len(graph.vert) // numero corrente di nodi presenti nel grafo
	insertedNode := 0            // nodi inseriti in questa chiamata
	if node == adjacent {        // controllo per far si che un nodo non possa essere vicino di se stesso
		return 0
	}
	if currVerst < graph.nodes { // se il numero di nodi presenti nel grafo è minore del numero massimo procedi
		if graph.vert[node] == nil { // se il nodo non esiste crealo e metterlo nel grafo
			graph.vert[node] = new(Vertex)
			graph.vert[node].Key = node
			insertedNode++ // incremento numero di nodi aggiunti in questa chiamata
		}
		graph.vert[node].Data = data                       // Inserimento dati del nodo
		if graph.vert[adjacent] == nil && adjacent != "" { // se il nodo vicino di 'node' non è presente nel grafo aggiungilo
			graph.vert[adjacent] = new(Vertex)
			graph.vert[adjacent].Key = adjacent
			insertedNode++ // incremento numero di nodi aggiunti in questa chiamata
		}
		// aggiungi il nodo 'adjacent' alla lista di adiacenza di 'node'
		graph.Adj[*graph.vert[node]] = append(graph.Adj[*graph.vert[node]], graph.vert[adjacent])
	} else { // altrimenti
		if graph.vert[node] != nil && graph.vert[adjacent] != nil { // verifica che il nodo 'node' e 'adjacent' esistano già
			// se si aggiungi 'adjacent' alla lista di adiacenza di 'node', in questo caso verranno connessi solo nodi già
			// presenti nel grafo, e non verranno aggiunti nodi nuovi
			graph.Adj[*graph.vert[node]] = append(graph.Adj[*graph.vert[node]], graph.vert[adjacent])
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
	nomeNodo dato1 dato2 dato3:vicino

	i dati del nodo non sono necessari, quindo possono essere omessi, e di default il loro valore sarà una stringa vuota
*/

func ReadGraph(graph *Graph) {
	fmt.Println("Inserisci i nodi del grafo con i suoi nodi adiacenti nel seguente formato (per interrompere inserire -1): ")
	fmt.Println("nome dato1 dato2 dato3:adiacente")

	scanner := bufio.NewScanner(os.Stdin)

	for i := 0; i <= graph.nodes; {
		scanner.Scan()
		input := scanner.Text()
		if input == "-1" {
			break
		}
		data := strings.Split(input, ":")
		adjacent := strings.TrimSpace(strings.Split(data[1], " ")[0])

		node := strings.Split(data[0], " ")
		nodeKey := node[0]
		var nodeData [3]string
		if len(node) == 4 {
			nodeData = [3]string{node[1], node[2], node[3]}
		}
		i = i + fillGraph(graph, nodeKey, nodeData, adjacent)
	}
}

/*
	Permette di stampare il grafo nel formato
	nodo1 => vicino1-1, vicino1-2 ...
	nodo2 => vicino2-1, vicino2-2 ...
	nodo3 => vicino3-1, vicino3-2 ...
	...
*/

func PrintGraph(graph *Graph) {
	for vertexKey, vertex := range graph.vert {
		fmt.Printf("%s => ", vertexKey)
		for _, adj := range graph.Adj[*vertex] {
			fmt.Printf("%s ", adj.Key)
		}
		fmt.Println()
	}
}
