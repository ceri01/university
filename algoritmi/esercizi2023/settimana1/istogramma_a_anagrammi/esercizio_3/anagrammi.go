package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var word1 string
	var word2 string
	var err error

	fmt.Println("Inserisci una parola.")
	_, err = fmt.Scan(&word1)
	if err != nil {
		fmt.Println("Si è verificato un errore nell'input")
		os.Exit(-1)
	}

	fmt.Println("Inserisci una seconda frase.")
	_, err = fmt.Scan(&word2)
	if err != nil {
		fmt.Println("Si è verificato un errore nell'input")
		os.Exit(-1)
	}

	if len(word1) == len(word2) {
		for _, r := range word1 {
			if !strings.Contains(word2, string(r)) {
				fmt.Println("Le due parole non sono l'una l'anagramma dell'altra.")
				os.Exit(0)
			}
		}
		fmt.Println(word2, " è l'anagramma di ", word1)
	} else {
		fmt.Println("Le due parole non sono l'una l'anagramma dell'altra.")
	}
}
