// Stampa un conto alla rovescia

package main

import "fmt"

func main() {
	var n int
	fmt.Println("Inserisci un numero")
	fmt.Scan(&n)

	if n < 0 {
		fmt.Println("inserisci un numero maggiore di 0")
	} else {
		for n >= 0 {
			fmt.Print(n, " ")
			n--
		}
	}
}