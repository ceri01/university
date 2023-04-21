package struttureDati

import (
	"fmt"
)

type OrderedLinkedList struct {
	head *ListNode
}

func Insert(list *OrderedLinkedList, val int) { // Tempo lineare => O(n)
	if list != nil {
		node := newNode(val)
		el := list.head
		var prev *ListNode = nil
		for el != nil && el.Key <= val {
			prev = el
			el = el.next
		}
		if prev != nil {
			prev.next = node
		} else {
			list.head = node
		}
		node.next = el
	}
}

/*func SearchByKey(list *OrderedLinkedList, key int) (bool, *ListNode) { // Tempo lineare => O(n)
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

func SearchByPosition(list *OrderedLinkedList, position int) (bool, *ListNode) { // Tempo lineare => O(n)
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

func RemoveByVal(list *OrderedLinkedList, key int) { // Tempo lineare => O(n)
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

func RemoveByposition(list *OrderedLinkedList, position int) { // Tempo lineare => O(n)
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
}*/

func PrintOrderedLinkedList(list *OrderedLinkedList) { // Tempo lineare => O(n)
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
