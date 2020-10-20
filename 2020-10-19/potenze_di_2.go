// Stampa le potenze di 2 <= n

package main

import "fmt"

func main() {
	var n int
	fmt.Println("Inserisci un numero")
	fmt.Scan(&n)
	potenza := 2
	i := 1
	if n >= 2 {
		fmt.Println("2 alla", i, "= ", potenza)
	}
	for potenza < n {
		i++
		potenza *= 2
		if potenza <= n {
			fmt.Println("2 alla", i, "= ", potenza)
		}
	}
}