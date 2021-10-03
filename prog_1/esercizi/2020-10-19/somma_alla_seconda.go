// Stampa la somma dei numeri da 1^2 fino ad n^2

package main

import "fmt"

func main() {
	var n, somma int
	fmt.Println("Inserisci un numero")
	fmt.Scan(&n)

	if n < 0 {
		for n < 0{
			somma += n * n
			n++
		}
		fmt.Println("La somma dei numeri da 1^2 a n^2 è", somma, "|il numero è stato traformato in positivo|")
	} else {
		for n > 0{
			somma += n * n
			n--
		}
		fmt.Println("La somma dei numeri da 1^2 a n^2 è", somma)
	}
}