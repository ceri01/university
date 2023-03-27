// Stabilire se un numero intero contiene qualche 0 nelle ultime 3 cifre

package main

import "fmt"

func main() {
	var number int
	fmt.Println("Inserisci un numero intero")
	fmt.Scan(&number)
	if number % 1000 == 0 {
		fmt.Print("Il numero ha le ultime 3 cifre uguali a 0")
	}
}