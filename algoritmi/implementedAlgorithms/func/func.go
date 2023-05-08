package _func

/*	Max e Min tra interi
	Costo = 1 confronto
	Tempo costante
	Spazio costante, non vengono create variabili
*/

func MaxInt(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func MinInt(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

/*	search dell'elemento minimo/massimo in un array
	Costo tempo = O(n) devo iterare su tutto l'array
	Costo spazio = O(1) creo due variabili
*/

func RicercaMin(arr []int) int {
	min := 0
	dim := len(arr)
	for i := 1; i < dim; i++ {
		if arr[min] > arr[i] {
			min = i
		}
	}
	return min
}

func RicercaMax(arr []int) int {
	max := 0
	dim := len(arr)
	for i := 1; i < dim; i++ {
		if arr[max] < arr[i] {
			max = i
		}
	}
	return max
}

func Swap(arr []int, pos1, pos2 int) {
	tmp := arr[pos1]
	arr[pos1] = arr[pos2]
	arr[pos2] = tmp
}
