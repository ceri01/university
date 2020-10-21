// Stabilire se il numero inserito contiene almeno un 3

package main

import "fmt"

func main() {
	var n int
	fmt.Println("Inserisci un numero")
	fmt.Scan(&n)
	for ;n > 0; {
		cifra := n % 10
		if cifra == 3 {
			fmt.Println("Il numero contiene almeno un tre")
			break
		}
		n /= 10
	}
	if n == 0 {
		fmt.Println("il numero non contiene alcun tre")
	}
}
