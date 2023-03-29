package main

import (
	"fmt"
	"os"
)

func main() {
	couples := 0
	var current int
	var prev int

	_, err := fmt.Scanf("%d", &current)
	if err != nil {
		fmt.Println("Errore inserimento.")
		os.Exit(-1)
	}

	for ; current != 0; {
		prev = current
		_, err := fmt.Scanf("%d", &current)
		if err != nil {
			fmt.Println("Errore inserimento.")
			os.Exit(-1)
		}

		if current == prev {
			couples++
		}
	}

	fmt.Println("Coppie numeri naturali consecutive: ", couples)
}
