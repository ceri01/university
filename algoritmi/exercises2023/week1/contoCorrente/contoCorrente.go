package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var balance, cost int
	var err error

	args := os.Args
	if len(args) > 1 {
		balance, err = strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Errore nella conversione dell'input, possibile inserimento errato")
		}
	} else {
		fmt.Println("Errore, parametro mancante.")
		os.Exit(-1)
	}

	fmt.Println("Inserisci le spese")

	for balance > 0 {
		_, err = fmt.Scanf("%d", &cost)
		balance = balance - cost
		fmt.Printf("Saldo => %d\n", balance)
	}
}
