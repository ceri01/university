// MCD con teorema di Euclide

package main

import "fmt"

func main() {
	var x, y, r int
	fmt.Println("Inserisci il primo numero")
	fmt.Scan(&x)
	fmt.Println("Inserisci il secondo numero")
	fmt.Scan(&y)
	for r = x % y; r != 0; {
		x = y
		y = r
		r = x % y
	}
	fmt.Println("L'MCD è", y)
}