package util

/*	Max e Min tra interi
	Costo = 1 confronto
	Tempo costante
	Spazio costante, non vengono create variabili
*/

func Max_int(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func Min_int(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

/*	ricerca dell'elemento minimo/massimo in un array
	Costo tempo = O(n) devo iterare su tutto l'array
	Costo spazio = O(1) creo due variabili
*/

func Ricerca_min(arr []int) int {
	min := 0
	dim := len(arr)
	for i := 1; i < dim; i++ {
		if arr[min] > arr[i] {
			min = i
		}
	}
	return min
}

func Ricerca_max(arr []int) int {
	max := 0
	dim := len(arr)
	for i := 1; i < dim; i++ {
		if arr[max] < arr[i] {
			max = i
		}
	}
	return max
}
