package withoutComparison

/*	b è il numero più grande presente in array
	Costo tempo:
		caso migliore: O(n + b)
		caso medio: O(n + b)
		caso peggiore: O(n + b)
	Costo spazio: θ(n + b)
*/

func IntegerSort(data []int, b int) {
	size := len(data)
	support := make([]int, b)
	for i := 0; i < size; i++ { // conto le occorrenze dei valori presenti nella slice e li salvo in support
		support[data[i]]++
	}
	ind := 0
	for i := 0; i < b; i++ {
		for j := 0; j < support[i]; j++ {
			data[ind] = i
			ind++
		}
	}
}
