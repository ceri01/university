package struttureDati

import (
	"fmt"
)

type ListNode struct {
	Key  int // In questo caso Key è la chiave, ma possiamo aggiungere altri campi
	next *ListNode
}

type LinkedList struct {
	head *ListNode
}

func newNode(val int) *ListNode {
	// non è necessario deferenziare i puntatori in Go, infatti non è necessario dichiarare var node = *listNode
	// e utilizzarlo come in C.
	node := new(ListNode)
	node.Key = val
	node.next = nil
	return node
	// node = &listNode{val, nil} => versione più breve delle tre righe sopra
}

func InsertHead(list *LinkedList, val int) { // Tempo costante => O(1)
	// Era possibile passare la lista normalmente e poi restituire la lista modificata, ma il chiamante avrebbe dovuto
	// fare un assegnamento. In questo modo basterà passare l'indirizzo della lista.
	node := newNode(val)
	node.next = list.head
	list.head = node
}

func InsertTail(list *LinkedList, val int) { // Tempo lineare => O(n)
	node := newNode(val)
	el := list.head
	for el.next != nil {
		el = el.next
	}
	el.next = node
}

func SearchByKey(list *LinkedList, key int) (bool, *ListNode) { // Tempo lineare => O(n)
	if list != nil {
		el := list.head
		for el != nil && el.Key != key {
			el = el.next
		}
		if el != nil {
			return true, el
		}
	}
	return false, nil
}

func SearchByPosition(list *LinkedList, position int) (bool, *ListNode) { // Tempo lineare => O(n)
	if list != nil && position >= 0 {
		el := list.head
		for i := 0; el != nil && i < position; i++ {
			el = el.next
		}
		if el != nil {
			return true, el
		}
	}
	return false, nil
}

func RemoveByVal(list *LinkedList, key int) { // Tempo lineare => O(n)
	if list != nil {
		var prev *ListNode = nil
		el := list.head
		for el != nil && el.Key != key {
			prev = el
			el = el.next
		}
		if el != nil {
			if prev == nil {
				list.head = el.next
			} else {
				prev.next = el.next
			}
		}
	}
}

func RemoveByposition(list *LinkedList, position int) { // Tempo lineare => O(n)
	if list != nil && position > 0 {
		var prev *ListNode = nil
		el := list.head
		for i := 0; el != nil && i < position; i++ {
			prev = el
			el = el.next
		}
		if el != nil {
			if prev == nil {
				list.head = el.next
			} else {
				prev.next = el.next
			}
		}
	}
}

func PrintList(list *LinkedList) { // Tempo lineare => O(n)
	if list.head != nil {
		el := list.head
		fmt.Print("[")
		for el.next != nil {
			fmt.Printf("%d ", el.Key)
			el = el.next
		}
		fmt.Printf("%d]\n", el.Key)
	} else {
		fmt.Println("[]")
	}
}
