package ordinamento

func Merge_sort(arr []int) { // versione già migliorata
	// Costo tempo O(n*log(n))
	// Costo tempo O(log(n))
	dim := len(arr)
	arr2 := make([]int, dim)
	mergeSort(arr, 0, dim, arr2)
}

func mergeSort(arr []int, i int, f int, arr2 []int) {
	if f-i > 1 { // costo spazio O(log(n)), le due chiamate ricorsive occupano lo stesso spazio,
		// quando una chiamata termina la successiva occupa lo spazio gi' usato dall'altro record di attivazione
		m := (i + f) / 2
		mergeSort(arr, i, m, arr2)
		mergeSort(arr, m, f, arr2)
		merge(arr, i, m, f, arr2)
	}
}

func merge(arr []int, i int, m int, f int, arr2 []int) {
	i1 := i
	i2 := m
	k := 0

	for i1 < m && i2 < f { // Costo tempo O(n*log(n))
		if arr[i1] <= arr[i2] {
			arr2[k] = arr[i1]
			i1++
		} else {
			arr2[k] = arr[i2]
			i2++
		}
		k++
	}
	if i1 < m { // Costo tempo O(1)
		for ; i1 < m; i1++ {
			arr2[k] = arr[i1]
			k++
		}
	} else { // costo O(1)
		for ; i2 < f; i2++ {
			arr2[k] = arr[i2]
			k++
		}
	}

	for k := 0; k < f-i; k++ { // Costo tempo O(n)
		arr[i+k] = arr2[k]
	}
}
