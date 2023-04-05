package ricerca

/*	ricerca dicotomica iterativa in un array ORDINATO
	Costo tempo = O(lon(n))
	Costo spazio = O(1) creo quattro variabili
*/

func ricerca_dicotomica_iterativa(arr []int, toFind int) int {
	isx := 0        // indice sinistro
	idx := len(arr) // indice destro
	i := -1         // indice elemento da trovare
	var m int
	for isx < idx && i == -1 {
		// cerco fino a che l'indice sinistro è minore di quello destro, e fino a che i è uguale a -1
		m = (isx + idx) / 2   // ricavo indice medio
		if arr[m] == toFind { // se l'elemento da trovare è in posizione m, assegno a i il valore di m
			i = m
		} else if arr[m] < toFind {
			// se invece l'elemento in posizione m è minore dell'elemento da trovare assegno all'indice destro il
			// valore di m
			idx = m
		} else {
			// se invece l'elemento in posizione m è maggiore dell'elemento da trovare assegno all'indice sinistro il
			// valore di m + 1
			isx = m + 1
		}
	}
	return i
}
