// Stampa la somma delle cifre da 1 a n

package main

import "fmt"

func main() {
	var num, c, somma int
	fmt.Println("Inserisci un numero")
	fmt.Scan(&num)
	if num < 0 {
		fmt.Println("Inserisci un numero >= 0")
	} else {
		for c <= num{
			somma += c
			c++
		}
		fmt.Println("(Ciclo) somma delle cifre da 1 a n", num, "=", somma)
		somma = (num*(num+1))/2
		fmt.Println("(Formula) somma delle cifre da 1 a n", num, "=", somma)
	}
}