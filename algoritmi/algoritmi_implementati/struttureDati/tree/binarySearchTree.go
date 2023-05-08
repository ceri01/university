package tree

import (
	util "algoritmi/algoritmi/algoritmi_implementati/struttureDati"
	"algoritmi/algoritmi/algoritmi_implementati/struttureDati/queue"
	"fmt"
)

func CreateBinarySearchTree(val int) *util.BiSearchTree {
	tree := new(util.BiSearchTree)
	tree.Val = val
	tree.Left = nil
	tree.Right = nil
	return tree
}

func InsertRecursive(tree *util.BiSearchTree, val int) *util.BiSearchTree {
	if tree == nil {
		tree = CreateBinarySearchTree(val)
		tree.Left = nil
		tree.Right = nil
	} else if val < tree.Val {
		tree.Left = InsertRecursive(tree.Left, val)
	} else {
		tree.Right = InsertRecursive(tree.Right, val)
	}
	return tree
}

func InsertIterative(tree *util.BiSearchTree, val int) {
	newNode := CreateBinarySearchTree(val)
	curr := tree
	father := new(util.BiSearchTree)

	for curr != nil {
		father = curr
		if val < curr.Val {
			curr = curr.Left
		} else {
			curr = curr.Right
		}
	}

	if father == nil {
		tree = newNode
	} else if val < father.Val {
		father.Left = newNode
	} else {
		father.Right = newNode
	}
}

func FindRecursive(tree *util.BiSearchTree, toFind int) *util.BiSearchTree {
	if tree == nil {
		return nil
	} else if toFind > tree.Val {
		return FindRecursive(tree.Right, toFind)
	} else if toFind < tree.Val {
		return FindRecursive(tree.Left, toFind)
	}
	// Le due ricorsioni sopra sono ricorsioni in coda, di conseguenza è possibile sostituire le chiamate ricorsive
	// Con una semplice iterazione, in modo da ottimizzare la quantità di spazio occupato in memoria.
	// Il tempo impiegato sarà lo stesso in qunto la ricorsione in coda ha lo stesso consumo di tempo di una funzione ricorsiva
	return tree
}

func FindIterative(tree *util.BiSearchTree, toFind int) *util.BiSearchTree {
	tmp := tree
	for tmp != nil && tmp.Val != toFind {
		if toFind > tmp.Val {
			tmp = tmp.Right
		} else if toFind < tmp.Val {
			tmp = tmp.Left
		}
	}
	return tmp
}

func Remove(tree *util.BiSearchTree, val int) {
	tmp := tree
	sx := false
	father := tree

	for tmp != nil && tmp.Val != val {
		father = tmp
		if val > tmp.Val {
			sx = false
			tmp = tmp.Right
		} else if val < tmp.Val {
			sx = true
			tmp = tmp.Left
		}
	}

	if tmp.Right == nil && tmp.Left == nil { // Eliminaizone foglia
		if sx {
			father.Left = nil
		} else {
			father.Right = nil
		}
	} else if tmp.Right != nil && tmp.Left == nil {
		if sx {
			father.Left = tmp.Right
		} else {
			father.Right = tmp.Right
		}
	} else if tmp.Right == nil && tmp.Left != nil {
		if sx {
			father.Left = tmp.Left
		} else {
			father.Right = tmp.Left
		}
	} else if tmp.Left != nil && tmp.Right != nil {
		max := Max(tmp.Left)
		tmp.Val = max.Val
		max.Val = val
		Remove(tmp.Left, val)
		/*
			Sarebbe una ricorsione in coda, quindi sostituibile da un ciclo (iterazione) per risparmiare spazio nello heap,
			ma siccome questa chiamata verrà fatta al massimo una volta per eliminaizone, possiamo dire che lo spaizo usato dalla
			ricorsione è trascurabile.
		*/
	}
}

func Max(tree *util.BiSearchTree) *util.BiSearchTree {
	if tree != nil {
		curr := tree
		for curr.Right != nil {
			curr = curr.Right
		}
		return curr
	}
	return nil
}

func InOrderSearchBiTree(root *util.BiSearchTree) {
	// Tramite questa visita dell'albero binario di ricerca è possibile visualizzare le chiavi in ordine
	if root != nil {
		InOrderSearchBiTree(root.Left)
		fmt.Printf("%d ", root.Val)
		InOrderSearchBiTree(root.Right)
	}
}

func PreOrderSearchBiTree(root *util.BiSearchTree) {
	if root != nil {
		fmt.Printf("%d ", root.Val)
		PreOrderSearchBiTree(root.Left)
		PreOrderSearchBiTree(root.Right)
	}
}

func PostOrderSearchBiTree(root *util.BiSearchTree) {
	if root != nil {
		PostOrderSearchBiTree(root.Left)
		PostOrderSearchBiTree(root.Right)
		fmt.Printf("%d ", root.Val)
	}
}

func DFS(tree *util.BiSearchTree, mode func(root *util.BiSearchTree)) {
	mode(tree)
	fmt.Println()
}

func BFS(tree *util.BiSearchTree) {
	if tree != nil {
		nodeQueue := queue.CreateBiSearchTreeQueue()
		nodeQueue.Enqueue(tree)
		for !nodeQueue.IsEmpty() {
			node := nodeQueue.Dequeue()
			if node != nil {
				fmt.Printf("%d ", node.Val)
				nodeQueue.Enqueue(node.Left)
				nodeQueue.Enqueue(node.Right)
			}
		}
		fmt.Println()
	}
}

func SummaryView(tree *util.BiSearchTree, count int) {
	fmt.Printf("%*s*", count, "")
	if tree != nil {
		fmt.Printf("%d\n", tree.Val)
		if tree.Left != nil || tree.Right != nil {
			SummaryView(tree.Left, count+2)
			SummaryView(tree.Right, count+2)
		}
	} else {
		fmt.Println()
	}
}
