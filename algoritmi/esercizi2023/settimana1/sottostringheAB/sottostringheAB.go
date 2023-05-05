package main

import (
	"fmt"
	"os"
)

func main() {
	var str string
	var Acount int
	var substring int

	_, err := fmt.Scan(&str)
	if err != nil {
		fmt.Println("Errore inserimento")
		os.Exit(-1)
	}

	for _, ch := range str {
		if ch == 'a' {
			Acount++
		} else if ch == 'b' {
			substring += Acount
		}
	}

	fmt.Println("Numero di sottostringhe = ", substring)

}
