// Stampa un triangolo

package main

import "fmt"

func main() {
        var n int
        fmt.Println("Inserisci un numero")
        fmt.Scan(&n)
        for i := 0; i < n; i++ {
                for j := 0; j < n; j++ {
			if j < i {
				fmt.Print(" ")
			} else {
				fmt.Print("*")

			}
		}
                fmt.Println("\n")
        }
}

