package main

import (
	"elaborazione_stringhe/funzioni"
	"fmt"
	"os"
)

func main() {
	const N = 4
	var err error
	parole := make([]string, N)

	for i := 0; i < N; i++ {
		_, err = fmt.Scan(&parole[i])
		if err != nil {
			fmt.Println("Errore nell'input")
			os.Exit(-1)
		}
	}

	fmt.Printf("\n")
	fmt.Println(funzioni.QuanteConA(parole))
	fmt.Printf("\n")
	fmt.Println(funzioni.PrimaConA(parole))
	fmt.Printf("\n")
	fmt.Println(funzioni.PiuCorta(parole))
}
