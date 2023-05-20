package graph

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Vertex struct {
	Data [3]string
	Key  string
}

type Graph struct {
	vert map[string]*Vertex
	Adj  map[Vertex][]*Vertex
}

func NewGraph(verts int, vertexKeys []string, vertexData [][]string) *Graph {
	graph := new(Graph)
	graph.vert = make(map[string]*Vertex, verts)
	for i := 0; i < verts; i++ {
		vertex := new(Vertex)
		vertex.Key = vertexKeys[i]
		vertex.Data[0] = vertexData[i][0]
		vertex.Data[1] = vertexData[i][1]
		vertex.Data[2] = vertexData[i][2]
		graph.vert[vertex.Key] = vertex
		graph.Adj = make(map[Vertex][]*Vertex, 1)
	}
	return graph
}

func FillGraph(graph *Graph) {
	fmt.Println(graph)
	fmt.Printf("Inserisci ...\n")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), " ")
		graph.Adj[*graph.vert[data[0]]] = append(graph.Adj[*graph.vert[data[0]]], graph.vert[data[1]])
	}
}

func PrintGraph(graph Graph) {
	for vertex, vertices := range graph.Adj {
		fmt.Printf("%s => ", vertex.Key)
		for _, adj := range vertices {
			fmt.Printf("%s ", adj.Key)
		}
		fmt.Println()
	}
}
