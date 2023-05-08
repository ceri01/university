package main

import (
	"elaborazione_serie_interi/funzioni"
	"fmt"
	"os"
)

func main() {
	const N = 10
	var err error
	numeri := make([]int, N)

	for i := 0; i < N; i++ {
		_, err = fmt.Scan(&numeri[i])
		if err != nil {
			fmt.Println("Errore nell'input")
			os.Exit(-1)
		}
	}

	fmt.Printf("\n")
	fmt.Println(funzioni.StranoProdotto(numeri))
	fmt.Printf("\n")
	funzioni.PariDispari(numeri)
	fmt.Printf("\n")
	fmt.Println(funzioni.MinDispari(numeri))
}
