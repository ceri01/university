// Dato n genera i primi n numeri primi

package main

import "fmt"

func main() {
	var n, k int
	fmt.Println("Inserisci un numero")
	fmt.Scan(&n)
	for i := 2; k != n; {
		var divisore int
		for divisore = 2; divisore < i; divisore++ {
			if i % divisore == 0 {
				break
			}
		}
		if !(divisore > i) && (i / divisore) == 1 {
			fmt.Println("numero primo =", divisore)
			k++
		}
		i++
	}
}