// Stampa quante cifre ha il numeor dato

package main

import "fmt"

func main() {
	var n, c int
	fmt.Println("Inserisci un numero")
	fmt.Scan(&n)
	if n==0 {
		c = 1
	} else {
		for n != 0{
			n /= 10
			c++
		}
	}
	fmt.Println("Numero di cifre:", c)
}