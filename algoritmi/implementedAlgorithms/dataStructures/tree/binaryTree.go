package tree

import (
	util "algoritmi/algoritmi/implementedAlgorithms/dataStructures"
	"algoritmi/algoritmi/implementedAlgorithms/dataStructures/queue"
	"fmt"
)

type BiTree struct {
	Root *util.BiTreeNode
}

func CreateBinaryTree() *BiTree {
	tree := new(BiTree)
	tree.Root = nil
	return tree
}

func InOrder(root *util.BiTreeNode) {
	if root != nil {
		InOrder(root.Left)
		fmt.Printf("%d ", root.Val)
		InOrder(root.Right)
	}
}

func PreOrder(root *util.BiTreeNode) {
	if root != nil {
		fmt.Printf("%d ", root.Val)
		PreOrder(root.Left)
		PreOrder(root.Right)
	}
}

func PostOrder(root *util.BiTreeNode) {
	if root != nil {
		PreOrder(root.Left)
		PreOrder(root.Right)
		fmt.Printf("%d ", root.Val)
	}
}

func (tree *BiTree) DFS(mode func(root *util.BiTreeNode)) {
	mode(tree.Root)
	fmt.Println()
}

func (tree *BiTree) BFS() {
	if tree.Root != nil {
		nodeQueue := queue.CreateTreeNodeQueue()
		nodeQueue.Enqueue(tree.Root)
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

func (tree *BiTree) Nnode(node *util.BiTreeNode) int {
	if node == nil {
		return 0
	} else {
		nsx := tree.Nnode(node.Left)
		ndx := tree.Nnode(node.Right)
		return 1 + nsx + ndx
	}
}
