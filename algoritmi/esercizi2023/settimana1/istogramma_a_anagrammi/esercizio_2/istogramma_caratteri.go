package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

const ALower = 'a'
const ZLower = 'z'

func main() {
	var phrase string
	chars := make([]int, 26, 26)

	fmt.Println("Inserisci una frase.")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	phrase = scanner.Text()

	if scanner.Err() != nil {
		fmt.Println("Si è verificato un errore nell'inserimento")
		os.Exit(-1)
	}

	for _, rn := range phrase {
		lowerCh := unicode.ToLower(rn)
		if lowerCh >= ALower && lowerCh <= ZLower {
			//fmt.Println(" rune in lower => ", lowerCh, "first rune => ", ALower)
			chars[lowerCh-ALower]++
		}
	}

	for key, occ := range chars {
		if occ > 0 {
			fmt.Println(string(rune(ALower+key)), " ", strings.Repeat("*", occ))
		}
	}
}
