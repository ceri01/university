// Stabilire se la somma delle cifre di un numero a tre cifre è > 10

package main

import (
	"fmt"
)

func main() {
	var n, prima, seconda, terza, somma int
	fmt.Println("Inserisci un numero")
	fmt.Scan(&n)
	terza = n % 10
	prima = (n - (n % 100)) / 100
	seconda = ((n % 100) - terza) / 10
	if prima != 0 {
		fmt.Println("Inserisci il numero senza lo 0 davanti")
	} else if n > 999 {
		fmt.Println("Inserisci il numero con massimo 3 cifre")
	} else {
		fmt.Println(prima)
		fmt.Println(seconda)
		fmt.Println(terza)
		somma = prima + seconda + terza
		if somma > 10 {
			fmt.Println("La somma delle 3 cifre è", somma, "ed è maggiore di 10")
		} else {
			fmt.Println("La somma delle 3 cifre è", somma, "e non è maggiore di 10")
		}
	}
}

