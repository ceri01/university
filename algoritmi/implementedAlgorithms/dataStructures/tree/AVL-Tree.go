package tree

import (
	util "algoritmi/algoritmi/implementedAlgorithms/dataStructures"
	"algoritmi/algoritmi/implementedAlgorithms/dataStructures/queue"
	"fmt"
)

func CreateAVLTree(val int) *util.AVLTree {
	tree := new(util.AVLTree)
	tree.Val = val
	tree.Left = nil
	tree.Right = nil
	tree.Height = 1
	return tree
}

func balanceFactor(tree *util.AVLTree) int {
	if tree != nil {
		return Height(tree.Left) - Height(tree.Right)
	}
	return 0
}

func Height(tree *util.AVLTree) int {
	if tree != nil {
		return tree.Height
	}
	return 0
}

func rotateLeft(tree *util.AVLTree) *util.AVLTree {
	treeDx := tree.Right
	tree.Right = treeDx.Left
	treeDx.Left = tree
	tree.Height = util.MaxInt(Height(tree.Left), Height(tree.Right)) + 1
	treeDx.Height = util.MaxInt(Height(tree.Left), Height(tree.Right)) + 1
	return treeDx
}

func rotateRight(tree *util.AVLTree) *util.AVLTree {
	treeSx := tree.Left
	tree.Left = treeSx.Right
	treeSx.Right = tree
	tree.Height = util.MaxInt(Height(tree.Left), Height(tree.Right)) + 1
	treeSx.Height = util.MaxInt(Height(tree.Left), Height(tree.Right)) + 1
	return treeSx
}

func Insert(tree *util.AVLTree, val int) *util.AVLTree {
	if tree == nil {
		return CreateAVLTree(val)
	} else if val < tree.Val {
		tree.Left = Insert(tree.Left, val)
	} else if val > tree.Val {
		tree.Right = Insert(tree.Right, val)
	} else {
		return tree
	}

	tree.Height = 1 + util.MaxInt(Height(tree.Left), Height(tree.Right))
	bFactor := balanceFactor(tree)

	if bFactor > 1 {
		fmt.Println("sbilanciato a sx ")
		if val < tree.Left.Val {
			return rotateRight(tree)
		} else if val > tree.Left.Val {
			tree.Left = rotateLeft(tree.Left)
			rotateRight(tree)
		}
	}
	if bFactor < -1 {
		fmt.Println("sbilanciato a dx ")
		if val > tree.Right.Val {
			return rotateLeft(tree)
		} else if val < tree.Right.Val {
			tree.Left = rotateRight(tree.Right)
			rotateLeft(tree)
		}
	}
	return tree
}

func Find(tree *util.BiSearchTree, toFind int) *util.BiSearchTree {
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

func RemoveNode(tree *util.AVLTree, val int) *util.AVLTree {
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
		max := MaxInAVL(tmp.Left)
		tmp.Val = max.Val
		max.Val = val
		RemoveNode(tmp.Left, val)
		/*
			Sarebbe una ricorsione in coda, quindi sostituibile da un ciclo (iterazione) per risparmiare spazio nello heap,
			ma siccome questa chiamata verrà fatta al massimo una volta per eliminaizone, possiamo dire che lo spaizo usato dalla
			ricorsione è trascurabile.
		*/
	}

	bFactor := balanceFactor(tree)

	if bFactor > 1 {
		if val < tree.Left.Val {
			return rotateRight(tree)
		} else if val > tree.Left.Val {
			tree.Left = rotateLeft(tree)
			rotateRight(tree)
		}
	}
	if bFactor < -1 {
		if val < tree.Right.Val {
			return rotateLeft(tree)
		} else if val > tree.Right.Val {
			tree.Left = rotateRight(tree)
			rotateLeft(tree)
		}
	}
	return tree
}

func MinInAVL(tree *util.AVLTree) *util.AVLTree {
	if tree != nil {
		curr := tree
		for curr.Left != nil {
			curr = curr.Left
		}
		return curr
	}
	return nil
}

func MaxInAVL(tree *util.AVLTree) *util.AVLTree {
	if tree != nil {
		curr := tree
		for curr.Right != nil {
			curr = curr.Right
		}
		return curr
	}
	return nil
}

func InOrderAVL(root *util.AVLTree) {
	// Tramite questa visita dell'albero binario di search è possibile visualizzare le chiavi in ordine
	if root != nil {
		InOrderAVL(root.Left)
		fmt.Printf("%d ", root.Val)
		InOrderAVL(root.Right)
	}
}

func PreOrderAVL(root *util.AVLTree) {
	if root != nil {
		fmt.Printf("%d ", root.Val)
		PreOrderAVL(root.Left)
		PreOrderAVL(root.Right)
	}
}

func PostOrderAVL(root *util.AVLTree) {
	if root != nil {
		PostOrderAVL(root.Left)
		PostOrderAVL(root.Right)
		fmt.Printf("%d ", root.Val)
	}
}

func DFSInAVL(tree *util.AVLTree, mode func(root *util.AVLTree)) {
	mode(tree)
	fmt.Println()
}

func BFSInAVL(tree *util.AVLTree) {
	if tree != nil {
		nodeQueue := queue.CreateAVLTreeQueue()
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

func SummaryViewAVL(tree *util.AVLTree, count int) {
	fmt.Printf("%*s*", count, "")
	if tree != nil {
		fmt.Printf("%d\n", tree.Val)
		if tree.Left != nil || tree.Right != nil {
			SummaryViewAVL(tree.Left, count+2)
			SummaryViewAVL(tree.Right, count+2)
		}
	} else {
		fmt.Println()
	}
}
