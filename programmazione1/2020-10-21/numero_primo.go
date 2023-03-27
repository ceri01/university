// Stabilire se un numero è primo

package main

import "fmt"

func main() {
	var n, i int
	fmt.Print("Inserisci un numero ")
	fmt.Scan(&n)
	for i = 2; i < n; i++ {
		if n % i == 0 {
			break
		}
	}
	if i < n {
		fmt.Println("Non è primo")
	} else {
		fmt.Println("Numero primo")
	}
}
