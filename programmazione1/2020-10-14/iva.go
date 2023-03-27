package main

import "fmt"

func main() {
	var prezzo, aliquota, imponibile float64
	fmt.Print("Inserisci il prezzo ")
	fmt.Scan(&prezzo)
	fmt.Print("Inserisci l'aliquota ")
	fmt.Scan(&aliquota)
	imponibile = prezzo * 100 / (100 + aliquota)
	fmt.Println("imponibile = ", imponibile)
}
