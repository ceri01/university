package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var str string
	var ch rune
	var err error

	fmt.Println("Inserisci una frase.")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err = scanner.Err()
	if err != nil {
		fmt.Println("Errore di inserimeto")
		os.Exit(-1)
	}
	str = scanner.Text()

	fmt.Println("inserisci il carattere i cui contare le occorrenze.")
	_, err = fmt.Scanf("%c", &ch)
	if err != nil {
		fmt.Println("Errore di inserimeto")
		os.Exit(-1)
	}

	fmt.Printf("La lettera %c compare %d volte.\n", ch, strings.Count(str, string(ch)))
}
