package queue

import (
	"fmt"
)

const DIM = 100

type ArrayQueue struct {
	begin int // Quando uno dei due indici "begin" o "end" arriveranno al valore massimo che possono assumere si verificherà un errore
	end   int // ma è un caso remoto che non ci interessa in quanto questa implementazione ha uno scopo didattico.
	empty bool
	full  bool
	array [DIM]int
}

func CreateArrayQueue() *ArrayQueue {
	queue := new(ArrayQueue)
	queue.begin = 0
	queue.end = -1
	queue.empty = true
	queue.full = false
	return queue
}

func (queue *ArrayQueue) IsEmpty() bool { // Costo tempo => costante O(1)
	return queue.empty
}

func (queue *ArrayQueue) Enqueue(val int) bool { // Costo tempo => costante O(1)
	// in questa variante viene ritornato un booleano per indicare se l'inserimento va a buon fine
	// non influisce eccessivamente sul tempo di esecuzione, e neanche sullo spazio occupato.
	if queue.full != true {
		queue.empty = false
		queue.end++
		queue.array[queue.end%DIM] = val
		if queue.empty != true && queue.end-queue.begin == DIM-1 {
			queue.full = true
		}
		return true
	}
	return false
}

func (queue *ArrayQueue) First() (int, bool) { // Costo tempo => costante O(1)
	// anche in questa variante vengono ritornati un booleano e un intero
	// il booleano indica se il valore restituito è da interpretare come dato oppure come
	// valore che specifica il fallimento dell'operaizone
	if queue.IsEmpty() {
		return -1, false
	}
	return queue.array[queue.begin%DIM], true
}

func (queue *ArrayQueue) Dequeue() (int, bool) { // Costo tempo => costante O(1)
	if queue.empty != true {
		queue.full = false
		res := queue.array[queue.begin%DIM]
		queue.begin++
		if queue.end-queue.begin < 0 {
			queue.empty = true
		}
		return res, true
	}
	return 0, false
}

func (queue *ArrayQueue) PrintQueue() { // Costo tempo => lineare O(n)
	i := queue.begin % DIM
	if queue != nil {
		if queue.IsEmpty() {
			fmt.Println("[]")
			return
		} else {
			fmt.Print("[")
			if queue.begin%DIM <= queue.end%DIM {
				for i = queue.begin % DIM; i < queue.end%DIM; i++ {
					fmt.Printf("%d ", queue.array[i])
				}
			} else {
				for i = queue.begin % DIM; i <= queue.end%DIM; i++ {
					if i >= DIM {
						i = -1
					} else {
						fmt.Printf("%d ", queue.array[i])
					}
				}
			}
			fmt.Printf("%d", queue.array[i])
			fmt.Print("]\n")
		}
	}
}
