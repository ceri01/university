package main

import (
	"fmt"
	"os"
)

func main() {
	var g int
	var val int

	fmt.Println("Inserisci numero di giorni da controllare ")

	_, err := fmt.Scanf("%d", &g)
	if err != nil {
		fmt.Println("valore inserito non valido.")
		os.Exit(-1)
	}

	for g < 3 { // controllo numero giorni minore di 3
		fmt.Println("inserire un valore maggiore di 3")
		_, err := fmt.Scanf("%d", &g)
		if err != nil {
			fmt.Println("valore inserito non valido.")
			os.Exit(-1)
		}
	}

	vals := make([]int, g, g)

	for i := 0; i < g; i++ { // inserimento valori
		fmt.Println("Inserisci consumo giorno ", i+1)
		_, err := fmt.Scanf("%d", &val)
		if err != nil {
			fmt.Println("valore inserito non valido.")
			os.Exit(-1)
		}
		vals[i] = val
	}

	if vals[0] < vals[1] {
		fmt.Println("consumo inferiore giorno => ", 1)
	}
	if vals[g-2] > vals[g-1] {
		fmt.Println("consumo inferiore giorno => ", g)
	}

	for i := 1; i < len(vals)-1; i++ {
		if vals[i] < vals[i-1] && vals[i] < vals[i+1] {
			fmt.Println("consumo inferiore giorno => ", i+1)
		}
	}
}
