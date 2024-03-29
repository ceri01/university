package orderedLinkedList

import (
	util "algoritmi/algoritmi/implementedAlgorithms/dataStructures"
	"fmt"
)

type OrderedLinkedList struct {
	head *util.ListNode
}

func (list *OrderedLinkedList) Insert(val int) { // Tempo lineare => O(n)
	if list != nil {
		node := util.NewNode(val)
		el := list.head
		var prev *util.ListNode = nil
		for el != nil && el.Key <= val {
			prev = el
			el = el.Next
		}
		if prev != nil {
			prev.Next = node
		} else {
			list.head = node
		}
		node.Next = el
	}
}

func (list *OrderedLinkedList) SearchByKey(key int) (bool, *util.ListNode) { // Tempo lineare => O(n)
	if list != nil {
		el := list.head
		for el != nil && el.Key < key {
			el = el.Next
		}
		if el != nil && el.Key == key {
			return true, el
		}
	}
	return false, nil
}

func (list *OrderedLinkedList) SearchByPosition(position int) (bool, *util.ListNode) { // Tempo lineare => O(n)
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

func (list *OrderedLinkedList) RemoveByVal(key int) { // Tempo lineare => O(n)
	if list != nil {
		var prev *util.ListNode = nil
		el := list.head
		for el != nil && el.Key < key {
			prev = el
			el = el.Next
		}
		if el != nil && el.Key == key {
			if prev == nil {
				list.head = el.Next
			} else {
				prev.Next = el.Next
			}
		}
	}
}

func (list *OrderedLinkedList) RemoveByPosition(position int) { // Tempo lineare => O(n)
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

func (list *OrderedLinkedList) PrintList() { // Tempo lineare => O(n)
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
