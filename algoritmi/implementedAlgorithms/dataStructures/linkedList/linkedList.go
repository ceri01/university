package linkedList

import (
	util "algoritmi/algoritmi/implementedAlgorithms/dataStructures"
	"fmt"
)

type LinkedList struct {
	head *util.ListNode
}

func (list *LinkedList) InsertHead(val int) { // Tempo costante => O(1)
	// Era possibile passare la lista normalmente e poi restituire la lista modificata, ma il chiamante avrebbe dovuto
	// fare un assegnamento. In questo modo basterà passare l'indirizzo della lista.
	node := util.NewNode(val)
	node.Next = list.head
	list.head = node
}

func (list *LinkedList) InsertTail(val int) { // Tempo lineare => O(n)
	node := util.NewNode(val)
	el := list.head
	for el.Next != nil {
		el = el.Next
	}
	el.Next = node
}

func (list *LinkedList) SearchByKey(key int) (bool, *util.ListNode) { // Tempo lineare => O(n)
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

func (list *LinkedList) SearchByPosition(position int) (bool, *util.ListNode) { // Tempo lineare => O(n)
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

func (list *LinkedList) RemoveByVal(key int) { // Tempo lineare => O(n)
	if list != nil {
		var prev *util.ListNode = nil
		el := list.head
		for el != nil && el.Key != key {
			prev = el
			el = el.Next
		}
		if el != nil {
			if prev == nil {
				list.head = el.Next
			} else {
				prev.Next = el.Next
			}
		}
	}
}

func (list *LinkedList) RemoveByPosition(position int) { // Tempo lineare => O(n)
	if list != nil && position >= 0 {
		var prev *util.ListNode = nil
		el := list.head
		for i := 0; el != nil && i < position; i++ {
			prev = el
			el = el.Next
		}
		if el != nil {
			if prev == nil {
				list.head = el.Next
			} else {
				prev.Next = el.Next
			}
		}
	}
}

func (list *LinkedList) PrintList() { // Tempo lineare => O(n)
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
