package ordinamento

import "algoritmi/algoritmi/algoritmi_implementati/func"

func Bubble_sort(arr []int) {
	i := 1
	dim := len(arr)
	swtch := true
	for swtch && i < dim {
		swtch = false
		for j := 1; j < dim-i; j++ { // oppure i <= dim - 1
			if arr[j] < arr[j-1] {
				_func.Swap(arr, j-1, j)
				swtch = true
			}
		}
		i++
	}
}
