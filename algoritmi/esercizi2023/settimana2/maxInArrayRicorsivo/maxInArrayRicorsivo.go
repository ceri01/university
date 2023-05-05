package main

import (
	"fmt"
	"os"
)

func max(val1, val2 int) int {
	if val1 >= val2 {
		return val1
	}
	return val2
}

func largest(numbers []int) int {
	n := len(numbers)
	if n == 1 {
		return numbers[0]
	} else {
		return max(numbers[n-1], largest(numbers[:n-1]))
	}
}

func main() {
	var dim, val int
	var arr []int
	fmt.Println("Inserisci dimensione array")
	_, err := fmt.Scanf("%d", &dim)
	if err != nil {
		fmt.Println("Errore di inserimento")
		os.Exit(-1)
	}

	for i := 0; i < dim; i++ {
		fmt.Println("Inserisci un valore")
		_, err := fmt.Scanf("%d", &val)
		if err != nil {
			fmt.Println("Errore di inserimento")
			os.Exit(-1)
		}
		arr = append(arr, val)
	}

	fmt.Println("Risultato =", largest(arr))
}
