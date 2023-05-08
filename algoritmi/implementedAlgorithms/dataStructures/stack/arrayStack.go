package stack

import "fmt"

type ArrayStack struct {
	top   int
	array [100]int
}

func CreateArrayStack() *ArrayStack {
	stack := new(ArrayStack)
	stack.top = -1
	return stack
}

func (stack *ArrayStack) IsEmpty() bool { // Costo tempo => costante O(1)
	return stack.top == -1
}

func (stack *ArrayStack) Push(val int) { // Costo tempo => costante O(1)
	stack.top++
	stack.array[stack.top] = val
}

func (stack *ArrayStack) Top() int { // Costo tempo => costante O(1)
	return stack.array[stack.top]
}

func (stack *ArrayStack) Pop() int { // Costo tempo => costante O(1)
	val := stack.array[stack.top]
	stack.top--
	return val
}

func (stack *ArrayStack) PrintStack() { // Costo tempo => lineare O(n)
	if stack != nil {
		fmt.Print("[")
		for i := 0; i <= stack.top-1; i++ {
			fmt.Printf("%d ", stack.array[i])
		}
		if !stack.IsEmpty() {
			fmt.Printf("%d]\n", stack.array[stack.top])
		} else {
			fmt.Print("]\n")
		}
	}
}
