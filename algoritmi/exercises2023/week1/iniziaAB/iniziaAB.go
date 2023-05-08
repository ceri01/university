package main

import (
	"fmt"
	"os"
)

func main() {
	var word string
	a, b := 0, 0
	for i := 0; i < 10; i++ {
		_, err := fmt.Scan(&word)
		if err != nil {
			fmt.Println("Errore nell'input.")
			os.Exit(-1)
		}
		if word[0] == 'a' {
			a++
		} else if word[0] == 'b' {
			b++
		}
	}
	fmt.Printf("Stringhe che iniziano con la lettera 'a' => %d.\n", a)
	fmt.Printf("Stringhe che iniziano con la lettera 'b' => %d.\n", b)
}
