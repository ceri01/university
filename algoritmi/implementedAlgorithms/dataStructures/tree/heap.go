package tree

import (
	_func "algoritmi/algoritmi/implementedAlgorithms/func"
)

func FixHeap(heap []int, r int, l int) {
	// r indice radice heap
	// l numero di elementi nello heap (ancora da ordinare)
	v := r           // posizione del dato da far scendere
	x := heap[v]     // dato da far scendere nella posizione appropriata
	toPlace := false // flag che indica se il nodo è ancora da collocare nella posizione corretta
	for {
		if (2*v)+1 >= l { // Se il figlio più grande è maggiore del numero di elementi nello heap (ancora da ordinare) fermati
			toPlace = true
		} else {
			u := maxChild(heap, l, v) // trova indice foglia con valore maggiore
			if heap[u] > x {          // se la foglia con valore maggiore è maggiore del nodo da far scendere allora:
				heap[v] = heap[u] // metti il figlio maggiore al posto del nodo da far scendere
				v = u             // il nuovo indice del nodo da far scendere è l'indice del figlio con valore maggiore
			} else {
				toPlace = true // altrimenti fermati
			}
		}
		if toPlace { // se toPlace è true fermati
			break
		}
	}
	heap[v] = x // viene posizionato il nodo da far scendere nella posizione corretta
}

func maxChild(arr []int, size int, curr int) int {
	/*	Si verifica che la posizione data da curr moltiplicata per due, il tutto sommato a due sia minore della dimensione della slice
		(ovvero che il nodo preso in considerazione possa avere un figlio destro), inoltre, se ciò è vero, si verifica che il figlio destro contenga un valore
		superiore al figlio sinistro.
	*/
	if (2*curr)+2 < size && arr[(2*curr)+2] > arr[(2*curr)+1] {
		return (2 * curr) + 2 // se ciò è vero viene ritornata la posizione del figlio destro
	}
	return (2 * curr) + 1 // altrimenti viene ritornata la posizione del figlio sinistro
}

//	Questa funzione permette la creazione di un heap (sempre sotto forma di slice)

func BuildHeap(heap []int) {
	size := len(heap) // in Go len() costa O(1) quindi non influisce sulle prestazioni
	// partendo dalla dimensione della slice applico FixHeap size - 1 volte (partendo da size - 1, quindi dalle foglie, fino ad arrivare a zero)
	for curr := size - 1; curr >= 0; curr-- {
		FixHeap(heap, curr, size) // sistemo la slice per farlo diventare un heap
	}
}

/*
	Per applicare l'heapsort non è necessario utilizzare effettivamente un heap come struttura dati, è possibile sfruttare
	direttamente un slice, e in questo modo l'algoritmo è in loco, ovvero non usa strutture di appoggio aggiuntive

	l'heapsort non è stabile, ovvero l'ordine dei duplicati non è mantenuto
	Costi:
		Tempo: se i confronti sono costanti la complessità dell'heapsort è O(n*log(n))
		Spazio: Grazie all'implementazione tramite array (slice in questo caso) l'algoritmo è in loco, ovvero non usa
				strutture aggiuntive, per questo motivo lo spazio è costante O(1)
*/

func HeapSort(heap []int) {
	lenHeap := len(heap) // ricavo la lunghezza della slice
	/*	applico BuldHeap alla slice in ingresso in modo da renderlo un heap (non la struttura heap effettiva ma sistema la slice
		in modo tale che sia organizzato come la vista in ampiezza dell'heap vero)
	*/
	BuildHeap(heap)
	/*	A questo punto scambio la foglia più a destra (ultimo elemento della slice) con la radice, successivamente applico
		FixHeap alla slice. Questa operazione la faccio lenHeap - 1 volte (partendo da lenHeap fino ad arrivare a zero)
	*/
	for l := lenHeap - 1; l >= 0; l-- {
		_func.Swap(heap, 0, l)
		FixHeap(heap, 0, l)
	}
}
