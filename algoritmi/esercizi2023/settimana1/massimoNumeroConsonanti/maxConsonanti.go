package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var w string
	var times int

	fmt.Println("Inserisci il numero di parola da processare")
	_, err := fmt.Scan(&times)
	if err != nil {
		fmt.Println("Errore inserimento.")
		os.Exit(-1)
	}

	sl := make([]string, times, times)
	for i := 0; i < times; i++ {
		fmt.Println("Inserisci parola.")
		_, err := fmt.Scan(&w)
		if err != nil {
			fmt.Println("Errore inserimento.")
			os.Exit(-1)
		}
		sl[i] = w
	}
	fmt.Println("la parola con più consonanti ne ha ", maxConsonanti(sl))
}

func maxConsonanti(words []string) int {
	max := 0
	for _, word := range words {
		cp := word
		for _, vowel := range []string{"a", "e", "i", "o", "u"} {
			cp = strings.ReplaceAll(cp, vowel, "")

		}
		if max < len(cp) {
			max = len(cp)
		}
	}
	return max
}
