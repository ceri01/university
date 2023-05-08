package ordinamento

import "algoritmi/algoritmi/algoritmi_implementati/func"

func partitioning(arr []int, i int, f int) int {
	pivot := arr[i]
	sx := i
	dx := f
	for sx < dx {
		for arr[dx] > pivot {
			dx--
		}
		for sx < dx && arr[sx] <= pivot {
			sx++
		}
		if sx < dx {
			_func.Swap(arr, sx, dx)
		}
	}
	_func.Swap(arr, i, dx)
	return dx
}

// Versione senza ricorsione in coda (altezza stack costante)
// Inoltre questa versione riordina la partizione di array più piccola in modo da
// poter avere i record di attivazione più piccoli (la metà rispetto alla lunghezza dell'array originale)

func quickSort(arr []int, i int, f int) {
	for f-i > 1 {
		pivot := partitioning(arr, i, f)
		if pivot-i < f-pivot {
			quickSort(arr, i, pivot)
			i = pivot + 1
		} else {
			quickSort(arr, pivot+1, f)
			f = pivot
		}
	}
}

/*
	Costo tempo quicksort:
		Caso peggiore (raro) = O(n^2)
		Caso medio (1.39*n*log(n)) => molto vicino al caso migliore, di conseguenza rende rara
			la possibilità che is verifichi il caso peggiore
		Caso migliore (non significativo) = O(n*log(n))
	Costo spazio quickSort:
		Il numero di variabili usate è costante, quindi bisogna considerare lo spazio occupato dalla ricorsione (quindi
		dall'altezza dello stack, ricordando però che lo spazio occupato dalla prima ricorsione viene sostituito da
		quello occupato dalla seconda, senza sommarsi).

		Caso peggiore: θ(n) => caso in cui la suddivisione (partitioning) è sbilanciata in una delle due parti
		Caso migliore: θ(log(n))

		Questa situazione si ha quando si usa l'implementazione base, quella riportata in questo file è la versione
		migliorata, ovvero quella senza ricorsione in coda e facendo la ricorsione suppa parte più piccola (quindi si evitano i casi
		in cui l'array è ordinato o quasi ordinato).

		Di conseguenza il costo del caso peggiore in termini di spazio diventa:
		Caso peggiore θ(log(n))
*/

func QuickSort(arr []int) {
	dim := len(arr) - 1
	quickSort(arr, 0, dim)
}
