package funzioni

import "fmt"

func PariDispari(numeri []int) {
	for _, val := range numeri {
		if val%2 == 0 {
			fmt.Printf("%d è un numero pari\n", val)
		} else {
			fmt.Printf("%d è un numero dispari\n", val)
		}
	}
}
