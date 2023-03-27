// Stampa le prime n potenze di 2

package main

import "fmt"

func main() {
	var n int
	fmt.Println("Inserisci un numero")
	fmt.Scan(&n)
	potenza := 1

	for i := 1; n >= i; {
		potenza *= 2
		fmt.Println("2 alla", i, "= ", potenza)
		i++
	}
}