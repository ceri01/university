// Stampa le potenze di 2 <= n

package main

import "fmt"

func main() {
	var n int
	fmt.Println("Inserisci un numero")
	fmt.Scan(&n)
	potenza := 1
	i := 0
	for potenza <= n {
		fmt.Println("1 alla", i, "= ", potenza)
		i++
		potenza *= 2
	}
}
