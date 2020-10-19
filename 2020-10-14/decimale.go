package main

import "fmt"

func main() {
	var numero float64
	var decimale float64
	fmt.Print("Inserisci un numero decimale: ")
	fmt.Scan(&numero)
	decimale = numero - float64(int(numero))
	fmt.Println("parte decimale: ", decimale)
}
