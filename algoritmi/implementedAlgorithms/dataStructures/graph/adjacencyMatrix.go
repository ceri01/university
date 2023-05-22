package graph

import (
	"fmt"
)

/*
	Una seconda rappresentazione dei grafi (anche qui per grafi orientati) è quella della matrice di adiacenza,
	ovvero si ha una matrice n*n in cui su righe sono presenti gli archi uscenti e sulle colonne gli archi entranti.
	Anche in questa implementazione vi è il limite per cui i vettori devono avere un valore compreso tra zero e n (escluso),
	è necessaria un'implementazione più complessa per poter usare chiavi di valore qualsiasi (come nell'implementazione di
	mapAdjacencyList).
	Ogni elemento della matrice rappresenta un arco dal nodo indicato sulla riga al nodo indicato sulla colonna
	(questo per grafi orientati, mentre per i grafi non orientati la matrice sarà simmetrica).
	Il costo di questa rappresentazione in termini di spazio è nell'ordine di θ(n^2), perchè la matrice è n*n.
	La matrice d'incidenza è utile per fare rappresentazioni algebriche ma poco efficiente per quanto riguarda lo spazio.
*/

type AdjMatrixGraph struct {
	verts     int
	adjMatrix [][]int
}

/*
	Genera un grafo contenente 'verts' nodi vuoti
*/

func NewAdjMatrixGraph(verts int) *AdjMatrixGraph {
	graph := new(AdjMatrixGraph)
	graph.adjMatrix = make([][]int, verts)
	graph.verts = verts
	return graph
}

/*
	Permette di riempire il grafo con i dati passati, ritorna il numero di nodi adiacenti inseriti
*/

func fillAdjMatrixGraph(graph *AdjMatrixGraph, node int, adjacent int) int {
	if node == adjacent { // controllo per far si che un nodo non possa essere vicino di se stesso
		return 0
	} else if node >= graph.verts || adjacent >= graph.verts || node < 0 || adjacent < 0 {
		fmt.Println("Errore di inserimento, i nodi devono avere un valore minore di ", graph.verts)
		return 0
	}

	for i := len(graph.adjMatrix[node]); i < adjacent; i++ {
		graph.adjMatrix[node] = append(graph.adjMatrix[node], 0)
	}
	if len(graph.adjMatrix[node]) > adjacent {
		graph.adjMatrix[node][adjacent] = 1
	} else {
		graph.adjMatrix[node] = append(graph.adjMatrix[node], 1)
	}
	return 1
}

/*
	Permette di popolare il grafo tramite standard input.
	È necessario rispettare un formato per poter inserire i dati, se ciò non avviene non è garantito che la funzione
	operi come dovrebbe

	FORMATO:
	nodo:vicino
*/

func ReadAdjMatrixGraph(graph *AdjMatrixGraph) {
	fmt.Println("Inserisci i nodi del grafo con i suoi nodi adiacenti nel seguente formato (per interrompere basta inserire una qualsiasi lettera): ")
	fmt.Println("nodo:adiacente")

	var node, adjacent int

	for i := 0; i <= graph.verts*graph.verts; {
		_, err := fmt.Scanf("%d:%d", &node, &adjacent)
		if err != nil {
			break
		}
		i += fillAdjMatrixGraph(graph, node, adjacent)
	}
}

/*
	Permette di stampare il grafo nel formato
	nodo1 => vicino1-1, vicino1-2 ...
	nodo2 => vicino2-1, vicino2-2 ...
	nodo3 => vicino3-1, vicino3-2 ...
	...
*/

func PrintAdjMatrixGraph(graph *AdjMatrixGraph) {
	for node, adj := range graph.adjMatrix {
		fmt.Printf("%d => ", node)
		for val, isPresent := range adj {
			if isPresent != 0 {
				fmt.Printf("%d ", val)
			}
		}
		fmt.Println()
	}
}
