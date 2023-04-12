package main

import (
	"fmt"
	"os"
)

func rocks(height int) int {
	if height == 1 {
		return 1
	} else {
		return (height * height) + rocks(height-1)
	}
}

func main() {
	var val int
	fmt.Println("Inserisci l'altezza della piramide")
	_, err := fmt.Scanf("%d", &val)
	if err != nil {
		fmt.Println("Errore di inserimento")
		os.Exit(-1)
	}
	fmt.Println("La piramide è composta da: ", rocks(val))
}
