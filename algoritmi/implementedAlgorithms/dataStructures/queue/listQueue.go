package queue

import "fmt"

type ListQueue struct {
	list []int
}

func CreateListQueue() *ListQueue {
	queue := new(ListQueue)
	return queue
}

func (queue *ListQueue) IsEmpty() bool { // Costo tempo => costante O(1)
	return len(queue.list) == 0
}

func (queue *ListQueue) Enqueue(val int) {
	/*
		Costo tempo:
			Caso migliore => costante O(1)
			Caso peggiore => O(n)
			Questi costi sono dati dalla append che per come è implementata ha questa complessità
	*/
	queue.list = append(queue.list, val)
}

func (queue *ListQueue) First() (int, bool) { // Costo tempo => costante O(1)
	if !queue.IsEmpty() {
		return queue.list[0], true
	}
	return -1, false
}

func (queue *ListQueue) Dequeue() (int, bool) { // Costo tempo => costante O(1)
	if !queue.IsEmpty() {
		val := queue.list[0]
		queue.list = queue.list[1:]
		return val, true
	}
	return -1, false
}

func (queue *ListQueue) PrintQueue() { // Costo tempo => lineare O(n)
	if len(queue.list) <= 0 {
		fmt.Print("[]\n")
	} else {
		fmt.Print("[")
		for i := 0; i < len(queue.list)-1; i++ {
			fmt.Printf("%d ", queue.list[i])
		}
		fmt.Printf("%d]\n", queue.list[len(queue.list)-1])
	}
}
