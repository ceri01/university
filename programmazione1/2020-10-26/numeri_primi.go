// Dato un numero stampa i numeri primi fino ad n (compreso)

package main

import "fmt"

func main() {
	var n int
	fmt.Println("Inserisci un numero")
	fmt.Scan(&n)
	for i := 2; i <= n; i++ {
		var divisore int
		for divisore = 2; divisore < i; divisore++ {
			if i % divisore == 0 {
				break
			}
		}
		if !(divisore > i) && (i / divisore) == 1 {
			fmt.Println("numero primo =", divisore)
		}
	}
}
