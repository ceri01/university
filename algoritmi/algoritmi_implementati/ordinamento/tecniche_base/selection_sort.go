package ordinamento

import (
	"algoritmi/algoritmi/algoritmi_implementati/util"
)

/*	Selection sort
	Al passo n - 2 (il penultimo) abbiamo finito perchè l'elemento in posizione n - 1 sarà già
	nella posizione corretta

	costo tempo = θ(n^2)
	costo spazio = O(1)

*/

func Selection_sort(arr []int) {
	dim := len(arr)               // dimensione array passato
	for i := 0; i <= dim-2; i++ { // itera l'array per n-2 volte
		min := util.Ricerca_min(arr[i:]) // trova la posizione del minimo nel sotto-array arr[i:]
		// Questa funzione che ricerca il minimo fa n - k - 1 confronti, dove n è la dimensione di arr e k - 1 è l'ultima posizione dell'array
		// scambio l'elemeno in posizione i con il minimo (faccio m + i perchè la funzione Ricerca_min considera il sotto-array come un array che va da 0 a n-1,
		// quindi come se fosse un array qualsiasi, facendo min + i ricavo la posizione effettiva dell'elemento minimo)
		tmp := arr[i]
		arr[i] = arr[min+i]
		arr[min+i] = tmp
	}
}
