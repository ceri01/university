// Dato l'anno dire se si può votare alla camera, al senato, o no si può votare

package main

import "fmt"

func main() {
	var anno int
	fmt.Println("Inserisci l'anno di nascita")
	fmt.Scan(&anno)
	if anno >= 25 {
		fmt.Println("Puoi votare per la camera e per il senato")
	} else if anno >= 18 && anno < 25 {
		fmt.Println("Puoi votare solo per la camera")
	} else {
		fmt.Println("Non puoi ancora votare")
	}
}