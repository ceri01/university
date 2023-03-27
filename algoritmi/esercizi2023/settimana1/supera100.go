package main

import (
	"fmt"
	"os"
)

func main() {
	var i int
	num := 0

	fmt.Println("Inserisci dei numeri, inserisci -1 per terminare.")
	_, err := fmt.Scanf("%d", &i)

	for i != -1 {
		if err != nil {
			fmt.Println("Errore nell'input.")
			os.Exit(-1)
		}
		if num == 0 && i >= 100 {
			num = i
		}
		_, err = fmt.Scanf("%d", &i)
	}

	if num > 100 {
		fmt.Printf("Primo numero maggiore di 100 => %d.\n", num)
	} else {
		fmt.Println("Nessun numero maggiore di 100.")
	}
}
