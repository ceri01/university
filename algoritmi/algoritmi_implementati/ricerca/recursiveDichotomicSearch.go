package ricerca

/*	ricerca dicotomica ricorsiva in un array ORDINATO
	Costo tempo = O(lon(n))
	Costo spazio = 2 + O(log(n)) record di attivazione
*/

func recursiveDichotomicSearch(arr []int, toFind, isx, idx int) int {
	var m int       // dichiarazione m
	if idx <= isx { // se indice dx è minore o uguale dell'indice sx l'array è vuoto oppiure i parametri sono sbagliati
		return -1
	} else {
		m = (idx + isx) / 2   // trovo l'indice medio
		if toFind == arr[m] { // se l'elemento da trovare è in quella posizione ritorno m
			return m
		} else if toFind < arr[m] {
			// altrimenti se è minore dell'elemento in posizione m faccio una chiamata ricorsiva dove cerco nella
			// parte di array che va dalla prima posizione fino ad m
			return recursiveDichotomicSearch(arr, toFind, isx, m)
		} else {
			// altrimenti se è maggiore dell'elemento in posizione m faccio una chiamata ricorsiva dove cerco nella
			// parte di array che va dalla posizione m + 1 posizione fino alla posizione finale dell'array
			return recursiveDichotomicSearch(arr, toFind, m+1, idx)
		}
	}
}
