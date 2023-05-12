package tree

import (
	_func "algoritmi/algoritmi/implementedAlgorithms/func"
	"fmt"
)

func FixHeap(heap []int, r int, l int) {
	// r indice radice heap
	// l numero di elementi nello heap (ancora da ordinare)
	v := r       // posizione del dato da far scendere
	x := heap[v] // dato da far scendere nella posizione appropriata
	toPlace := false
	for {
		if (2*v)+1 >= l { // Se il figlio più grande è maggiore del numero di elementi nello heap (ancora da ordinare) fermati
			toPlace = true
		} else {
			u := maxChild(heap, l-1, v) // trova la posizione del figlio con valore maggiore
			if heap[u] > x {            //
				heap[v] = heap[u] // metti il figlio maggiore al posto della radice
				v = u
			} else {
				toPlace = true
			}
		}
		if toPlace {
			break
		}
	}
	heap[v] = x
}

func maxChild(arr []int, size int, curr int) int {
	if (2*curr)+2 < size && arr[(2*curr)+2] > arr[(2*curr)+1] {
		return (2 * curr) + 2
	}
	return (2 * curr) + 1
}

func BuildHeap(heap []int, size int) {
	for curr := size / 2; curr >= 0; curr-- {
		FixHeap(heap, curr, size)
	}
}

func HeapSort(heap []int) {
	lenHeap := len(heap)
	BuildHeap(heap, lenHeap-1)
	fmt.Println(heap)
	for l := lenHeap - 1; l >= 0; l-- {
		_func.Swap(heap, 0, l)
		FixHeap(heap, 0, l)
	}
}
