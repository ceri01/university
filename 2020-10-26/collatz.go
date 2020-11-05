// Dato un numero stampa la sua congettura di Collatz

package main

import "fmt"

func main() {
	var n int
	fmt.Print("Inserisci n: ")
	fmt.Scan(&n)
	if n == 0 {
		fmt.Println("Valore non consentito")
		return
	}
	for n != 1 {
		if n % 2 == 0 {
			n /= 2
			fmt.Println("n =", n)
		} else {
			n = (n * 3) + 1
			fmt.Println("n =", n)
		}
	}
}
