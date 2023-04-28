package stack

import (
	util "algoritmi/algoritmi/algoritmi_implementati/struttureDati"
	"algoritmi/algoritmi/algoritmi_implementati/struttureDati/linkedList"
	"fmt"
)

type CustomListStack struct {
	top  *util.ListNode
	list *linkedList.LinkedList
}

func CreateCustomListStack() *CustomListStack {
	stack := new(CustomListStack)
	stack.top = nil
	stack.list = new(linkedList.LinkedList)
	return stack
}

func (stack *CustomListStack) IsEmpty() bool { // Costo tempo => Costante O(1)
	return stack.top == nil
}

func (stack *CustomListStack) Top() int { // Costo tempo => Costante O(1)
	return stack.top.Key
}

func (stack *CustomListStack) Push(val int) { // Costo tempo => Costante O(1)
	stack.list.InsertHead(val)
	_, stack.top = stack.list.SearchByPosition(0)
	// in questo caso viene utilizzata una funzione con costo di tempo lineare, ma siccome viene cercato sempre l'elemento
	// in prima posizione il tempo di esecuzione è sempre lo stesso, quindi costante.
}

func (stack *CustomListStack) Pop() int { // Costo tempo => Costante O(1)
	el := stack.top.Key
	stack.list.RemoveByPosition(0)
	_, stack.top = stack.list.SearchByPosition(0)
	// Anche in questo caso vengono utilizzate delle funzioni con costo di tempo lineare, ma anche qui viene vincolato il caso migliore
	// che ha tempo costante, quindi possiamo dire che il costo di tempo della funzione Pop è costante
	return el
}

func (stack *CustomListStack) PrintStack() { // Costo tempo => lineare O(n)
	fmt.Print("[")
	if stack.top != nil {
		tmp := stack.top
		for tmp.Next != nil {
			fmt.Printf("%d ", tmp.Key)
			tmp = tmp.Next
		}
		fmt.Printf("%d", tmp.Key)
	}
	fmt.Printf("]\n")

}
