package main

import (
	"fmt"
	"os"
)

func mulRecursive(val1, val2 int) int {
	if val2 == 1 {
		return val1
	} else if val2 == 0 {
		return 0
	} else {
		return val1 + mulRecursive(val1, val2-1)
	}
}

func main() {
	var val1, val2 int
	fmt.Println("Inserisci primo numero")
	_, err := fmt.Scanf("%d", &val1)
	if err != nil {
		fmt.Println("Errore di inserimento")
		os.Exit(-1)
	}

	fmt.Println("Inserisci secondo numero")
	_, err = fmt.Scanf("%d", &val2)
	if err != nil {
		fmt.Println("Errore di inserimento")
		os.Exit(-1)
	}

	fmt.Println("risultato", mulRecursive(val1, val2))
}
