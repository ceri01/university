// Stampare a video la media dell'altezza di n persone

package main

import "fmt"

func main() {
	var n int
	fmt.Println("Inserisci il numero di persone")
	fmt.Scan(&n)
	somma := 0
	for i := 0; i < n; i++ {
		var altezza int
		fmt.Println("Inserisci altezza persona", i + 1)
		fmt.Scan(&altezza)
		somma += altezza
	}
	if !(n <= 0) {
		fmt.Println("La media delle altezze delle persone è", float64(somma) / float64(n))
	} else {
		fmt.Println("Ci deve essere almeno una persona")
	}
}
