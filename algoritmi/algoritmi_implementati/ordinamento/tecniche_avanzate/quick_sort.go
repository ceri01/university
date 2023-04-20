package ordinamento

import "algoritmi/algoritmi/algoritmi_implementati/util"

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
			util.Swap(arr, sx, dx)
		}
	}
	util.Swap(arr, i, dx)
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

func Quick_sort(arr []int) {
	dim := len(arr) - 1
	quickSort(arr, 0, dim)
}
