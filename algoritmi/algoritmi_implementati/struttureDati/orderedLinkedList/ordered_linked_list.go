package orderedLinkedList

import (
	util "algoritmi/algoritmi/algoritmi_implementati/struttureDati"
	"fmt"
)

type OrderedLinkedList struct {
	head *util.ListNode
}

func Insert(list *OrderedLinkedList, val int) { // Tempo lineare => O(n)
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

func SearchByKey(list *OrderedLinkedList, key int) (bool, *util.ListNode) { // Tempo lineare => O(n)
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

func SearchByPosition(list *OrderedLinkedList, position int) (bool, *util.ListNode) { // Tempo lineare => O(n)
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

func RemoveByVal(list *OrderedLinkedList, key int) { // Tempo lineare => O(n)
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

func RemoveByPosition(list *OrderedLinkedList, position int) { // Tempo lineare => O(n)
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

func PrintList(list *OrderedLinkedList) { // Tempo lineare => O(n)
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
