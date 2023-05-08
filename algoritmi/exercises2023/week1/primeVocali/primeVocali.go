package main

import (
	"fmt"
	"os"
)

func main() {
	var word string
	vowels := map[int32]rune{
		97:  'a',
		101: 'e',
		105: 'i',
		111: 'o',
		117: 'i',
	}
	val := false

	for {
		_, err := fmt.Scan(&word)
		if err != nil {
			fmt.Println("Errore nell'input")
			os.Exit(-1)
		}
		for _, ch := range word {
			_, isPresent := vowels[ch]
			if isPresent {
				fmt.Println("prima volale: ", string(ch))
				val = true
				break
			}
		}
		if !val {
			fmt.Println("La parola non contiene vocali")
		}
		val = false
	}
}
