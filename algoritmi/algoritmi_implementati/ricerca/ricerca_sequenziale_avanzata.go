package ricerca

/*	ricerca iterativa
	Ritorna la posizione dell'elemento da trovare
	Con questa versione che parte dal fondo si risparmia il costo della selezione finale

	Costo tempo = O(n)
	Costo spazio = O(1) -> vengono create due variabili
*/

func ricerca_iterativa(arr []int, toFind int) int {
	i := len(arr) - 1 // indice ultimo elemento array

	// itera fino a che non l'indice è maggiore o uguale a 0 e l'elemento in posizione i
	// è diverso dall'elemento da trovare
	for i >= 0 && arr[i] != toFind {
		i--
	}
	// In questo modo rtorna -1 se l'elemento non è presente, altrimenti torna l'indice dell'elemento da trovare
	return i
}
