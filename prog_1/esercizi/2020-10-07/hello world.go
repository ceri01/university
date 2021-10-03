package main

import "fmt"

func main() {
	var i, n int
	fmt.Println("Quante volte vuoi essere salutato?")
	fmt.Scan(&n)
	for i = 0; i < n; i++ {
		fmt.Println("ciao!")
	}
}
