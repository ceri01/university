// Stabilire se un anno è bisestile (il primo anno bisestile parte dall'8 d.c.)
package main

import "fmt"

func main() {
	var anno int
	fmt.Println("Inserisci anno")
	fmt.Scan(&anno)
	if anno >= 8 && anno % 4 == 0 {
		fmt.Println("Il", anno, "è un anno bisestile")
	} else if anno < 8 {
		fmt.Println("Gli anni bisestili iniziano dall'8 d.c. inserisci un numero >= 8")
	} else {
		fmt.Println("Il", anno, "non è un anno bisestile")
	}
}