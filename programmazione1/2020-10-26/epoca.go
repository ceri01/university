// inserita una data stampa la differenza di giorni tra la data inserita e il 1/1/1970

package main

import "fmt"

func main() {
	var g, m, a int
	fmt.Print("Giorno: ")
	fmt.Scan(&g)
	fmt.Print("Mese: ")
	fmt.Scan(&m)
	fmt.Print("Anno: ")
	fmt.Scan(&a)
	
	gg := 0
	for anno := 1970; anno != a; {
		gg += 365
		if anno % 4 == 0 {
			gg++
		}
		if anno > a {
			anno--
		} else {
			anno++
		}
	}

	for mese := 1; mese < m; mese++ {
		if mese == 11 || mese == 4 || mese == 6 || mese == 9 {
			gg += 30
		} else if mese == 2 {
			gg += 28
			if a % 4 == 0 {
				gg += 1
			}
		} else {
			gg += 31
		}
	}
	gg += g

	fmt.Println("I giorni trascorsi dal 1/1/1970 al", g, "/", m, "/", a, " sono", (gg - 1), "\n")
}
