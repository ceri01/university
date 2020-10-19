package main
import "fmt"

func main() {
	var ora, minuti, manca int
	fmt.Println("Inserisci ora")
	fmt.Scan(&ora)
	fmt.Println("Inserisci minuti")
	fmt.Scan(&minuti)
	manca = (24-ora)*60 + (60 - minuti)
	fmt.Println("Mancano", manca, "alla prossima mezanotte")



}


