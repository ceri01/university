package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	val := 0
	number := -1
	maxZero := 0

	for val != -1 {
		fmt.Println("Inserisci numero.")
		_, err := fmt.Scan(&val)
		if err != nil {
			fmt.Println("Errore inserimento")
			os.Exit(-1)
		}

		valStr := strconv.Itoa(val)
		zeri := strings.Count(valStr, "0")
		if zeri > maxZero {
			maxZero = zeri
			number = val
		}
	}

	if number != -1 {
		fmt.Println("Il numero con più zeri è", number, "e ne ha ", maxZero)
	} else {
		fmt.Println("Non ci sono numeri contenenti la cifra zero")
	}
}
