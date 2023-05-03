package queue

import (
	"algoritmi/algoritmi/algoritmi_implementati/struttureDati/linkedList"
)

type CustomListQueue struct {
	list *linkedList.DoubleLinkedList
	size int
}

func CreateCustomListQueue() *CustomListQueue {
	queue := new(CustomListQueue)
	queue.size = 0
	return queue
}

func (queue *CustomListQueue) IsEmpty() bool { // Costo tempo => costante O(1)
	return queue.size == 0
}

func (queue *CustomListQueue) Enqueue(val int) { // Costo tempo => costante O(1)
	queue.list.InsertTail(val) // Usando le double linked list l'inserimento in coda ha costo di tempo costante
	queue.size++
}

func (queue *CustomListQueue) First() (int, bool) { // Costo tempo => costante O(1)
	if !queue.IsEmpty() {
		isPresent, val := queue.list.SearchByPosition(0) // La ricerca in prima posizione ha tempo costante (il ciclo viene saltato)
		return val.Key, isPresent
	}
	return -1, false
}

func (queue *CustomListQueue) Dequeue() (int, bool) { // Costo tempo => costante O(1)
	if !queue.IsEmpty() {
		val, _ := queue.First()
		queue.list.RemoveByPosition(0) // La rimozione in prima posizione ha tempo costante (il ciclo viene saltato)
		queue.size--
		return val, true
	}
	return -1, false
}

func (queue *CustomListQueue) PrintQueue() { // Costo tempo => lineare O(n)
	queue.list.PrintList()
}
