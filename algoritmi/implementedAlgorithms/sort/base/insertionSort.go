package ordinamento

/*
	Costo tempo:
		Caso peggiore => quadratico θ(n^2)
		Caso migliore => lineare θ(n-1) nel caso in cui l'array è già ordinato

		Ci sono due ciclo uno all'interno dell'altro, quindi numero di confronti quadratico
	Costo spazio: costante θ(1)

	Algoritmo in loco.
	Stabile.
*/

func InsertionSort(arr []int) {
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
