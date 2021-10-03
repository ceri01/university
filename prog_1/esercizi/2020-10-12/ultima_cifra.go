package main

import "fmt"

func main() {
	var numero int
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&numero)
	numero = (numero - 3)
	if numero % 10 == 0  {
		fmt.Println("Il numero inserito ha come ultima cira il 3")
	} else {
		fmt.Println("Il numero inserito non ha come ultima cira il 3")
	}
}
