package main

import (
	"fmt"
	"os"
)

func main() {
	var limit int
	var ascent int
	var val int
	var prev int

	isAscent := false

	fmt.Println("Inserisci il numero di altezze della catena montuosa")
	_, err := fmt.Scanf("%d", &limit)
	if err != nil {
		fmt.Println("Errore inserimento")
	}

	fmt.Println("Inserisci i valori delle altitudini")
	for i := 0; i < limit; i++ {
		prev = val
		_, err = fmt.Scanf("%d", &val)
		if err != nil {
			fmt.Println("Errore inserimento")
			os.Exit(-1)
		}

		if !(i == 0) {
			if isAscent && prev > val {
				ascent++
			}
			if prev < val {
				isAscent = true
			} else {
				isAscent = false
			}
		}
	}

	if isAscent && prev < val {
		ascent++
	}

	fmt.Println("numero di salite: ", ascent)
}
