package main
import "fmt"

func main() {
	var giorno, mese, capodanno, manca int
	capodanno = 12 * 30 + 1
	fmt.Println("Inserisci giorno")
	fmt.Scan(&giorno)
	fmt.Println("Inserisci mese")
	fmt.Scan(&mese)
	manca = capodanno - (((mese - 1) * 30) + giorno)
	fmt.Println("Mancano", manca - 6, "giorni a natale")
	fmt.Println("Mancano", manca, "giorni a capodanno")
}
