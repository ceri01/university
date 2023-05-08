package stack

import "fmt"

type ListStack struct {
	top  int
	list []int
}

func CreateListStack() *ListStack {
	stack := new(ListStack)
	stack.top = -1
	return stack
}

func (stack *ListStack) IsEmpty() bool { // Costo tempo => costante O(1)
	return stack.top == -1
}

func (stack *ListStack) Push(val int) {
	/*
		Costo tempo:
			Caso migliore => costante O(1)
			Caso peggiore => O(n)
			Questi costi sono dati dalla append che per come è implementata ha questa complessità
	*/
	stack.top++
	stack.list = append(stack.list, val)
}

func (stack *ListStack) Top() int { // Costo tempo => costante O(1)
	return stack.list[stack.top]
}

func (stack *ListStack) Pop() int { // Costo tempo => costante O(1)
	val := stack.list[stack.top]
	stack.top--
	return val
}

func (stack *ListStack) PrintStack() { // Costo tempo => lineare O(n)
	if stack != nil {
		fmt.Print("[")
		for i := 0; i <= stack.top-1; i++ {
			fmt.Printf("%d ", stack.list[i])
		}
		if !stack.IsEmpty() {
			fmt.Printf("%d]\n", stack.list[stack.top])
		} else {
			fmt.Print("]\n")
		}
	}
}
