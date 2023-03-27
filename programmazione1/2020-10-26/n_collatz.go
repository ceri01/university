// Dato un numero stampa le n congetture di collatz

package main

import "fmt"

func main() {
    var n int
    fmt.Print("Inserisci n: ")
    fmt.Scan(&n)
	var count int
	for count = 2; count <=  n; count++ {
		k := count
		fmt.Println("sequenza ", (k - 1))
		for k != 1 {
            if k % 2 == 0 {
                	k /= 2
                	fmt.Println("c =", k)
            } else {
                	k = (k * 3) + 1
                	fmt.Println("c =", k)
            }
        }
	}
}
