package queue

import util "algoritmi/algoritmi/algoritmi_implementati/struttureDati"

type TreeNodeQueue struct {
	list []*util.BiTreeNode
}

func CreateTreeNodeQueue() *TreeNodeQueue {
	queue := new(TreeNodeQueue)
	return queue
}

func (treeNode *TreeNodeQueue) IsEmpty() bool { // Costo tempo => costante O(1)
	return len(treeNode.list) == 0
}

func (treeNode *TreeNodeQueue) Enqueue(val *util.BiTreeNode) {
	/*
		Costo tempo:
			Caso migliore => costante O(1)
			Caso peggiore => O(n)
			Questi costi sono dati dalla append che per come è implementata ha questa complessità
	*/
	treeNode.list = append(treeNode.list, val)
}

func (treeNode *TreeNodeQueue) First() *util.BiTreeNode { // Costo tempo => costante O(1)
	if !treeNode.IsEmpty() {
		return treeNode.list[0]
	}
	return nil
}

func (treeNode *TreeNodeQueue) Dequeue() *util.BiTreeNode { // Costo tempo => costante O(1)
	if !treeNode.IsEmpty() {
		val := treeNode.list[0]
		treeNode.list = treeNode.list[1:]
		return val
	}
	return nil
}
