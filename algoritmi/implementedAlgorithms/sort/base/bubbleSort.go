package ordinamento

import "algoritmi/algoritmi/implementedAlgorithms/func"

/*
	Costo tempo:
		Caso peggiore => quadratico θ(n^2)
		Caso migliore => lineare θ(n-1) nel caso in cui l'array è già ordinato

		Ci sono due ciclo uno all'interno dell'altro, quindi numero di confronti quadratico
	Costo spazio: costante θ(1)

	Algoritmo in loco.
	Stabile.
*/

func BubbleSort(arr []int) {
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
