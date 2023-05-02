package linkedList

import (
	util "algoritmi/algoritmi/algoritmi_implementati/struttureDati"
	"fmt"
)

type DoubleLinkedList struct {
	head *util.ListNodeDouble
	tail *util.ListNodeDouble
}

func (list *DoubleLinkedList) InsertHead(val int) { // Tempo costante => O(1)
	// Era possibile passare la lista normalmente e poi restituire la lista modificata, ma il chiamante avrebbe dovuto
	// fare un assegnamento. In questo modo basterà passare l'indirizzo della lista.
	node := util.NewDoubleNode(val)
	node.Next = list.head
	list.head = node
	if list.head.Next != nil {
		list.head.Next.Prev = list.head
	}
	if list.tail == nil {
		list.tail = node
	}
}

func (list *DoubleLinkedList) InsertTail(val int) { // Tempo lineare => O(1)
	node := util.NewDoubleNode(val)
	node.Prev = list.tail
	list.tail = node
	if list.tail.Prev != nil {
		list.tail.Prev.Next = list.tail
	}
	if list.head == nil {
		list.head = node
	}
}

func (list *DoubleLinkedList) SearchByKey(key int) (bool, *util.ListNodeDouble) { // Tempo lineare => O(n)
	if list != nil {
		el := list.head
		for el != nil && el.Key != key {
			el = el.Next
		}
		if el != nil {
			return true, el
		}
	}
	return false, nil
}

func (list *DoubleLinkedList) SearchByPosition(position int) (bool, *util.ListNodeDouble) { // Tempo lineare => O(n)
	if list != nil && position >= 0 {
		el := list.head
		for i := 0; el != nil && i < position; i++ {
			el = el.Next
		}
		if el != nil {
			return true, el
		}
	}
	return false, nil
}

func (list *DoubleLinkedList) RemoveByVal(key int) { // Tempo lineare => O(n)
	if list != nil {
		el := list.head
		for el.Key != key {
			el = el.Next
		}
		if el != nil {
			if el.Next != nil {
				if el.Prev != nil {
					el.Prev.Next = el.Next
					el.Next.Prev = el.Prev
				} else {
					list.head = el.Next
					list.head.Prev = nil
				}
			} else {
				if el.Prev != nil {
					list.tail = el.Prev
					list.tail.Next = nil
				} else {
					list.tail = nil
					list.head = nil
				}
			}
		}
	}
}

func (list *DoubleLinkedList) RemoveByPosition(position int) { // Tempo lineare => O(n)
	if list != nil && position >= 0 {
		el := list.head
		for i := 0; el != nil && i < position; i++ {
			el = el.Next
		}
		if el != nil {
			fmt.Println(el)
			if el.Next != nil {
				if el.Prev != nil {
					el.Prev.Next = el.Next
					el.Next.Prev = el.Prev
				} else {
					list.head = el.Next
					list.head.Prev = nil
				}
			} else {
				if el.Prev != nil {
					list.tail = el.Prev
					list.tail.Next = nil
				} else {
					list.tail = nil
					list.head = nil
				}
			}
		}
	}
}

func (list *DoubleLinkedList) PrintList() { // Tempo lineare => O(n)
	if list.head != nil {
		el := list.head
		fmt.Print("[")
		for el.Next != nil {
			fmt.Printf("%d ", el.Key)
			el = el.Next
		}
		fmt.Printf("%d]\n", el.Key)
	} else {
		fmt.Println("[]")
	}
}
