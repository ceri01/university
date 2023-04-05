package ricerca

/*	ricerca iterativa
	Ritorna la posizione dell'elemento da trovare

	Costo tempo = O(n)
	Costo spazio = O(1) -> vengono create due variabili
*/

func ricerca_iterativa_base(arr []int, toFind int) int {
	dim := len(arr) // dimensione array
	i := 0          // indice

	// itera fino a che non l'indice è minore della dimensione e l'elemento in posizione i
	// è diverso dall'elemento da trovare
	for i < dim && arr[i] != toFind {
		i++
	}
	// se l'indice è uguale alla dimensione allora l'elemento non è presente nell'array
	if i == dim {
		return -1
	}
	// altrimenti ritorna l'indice
	return i
}
