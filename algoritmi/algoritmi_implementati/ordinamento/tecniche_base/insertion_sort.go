package ordinamento

func Insertion_sort(arr []int) {
	dim := len(arr)
	for i := 1; i <= dim-1; i++ {
		j := i - 1
		x := arr[i]
		for j >= 0 && arr[j] > x {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = x
	}
}
