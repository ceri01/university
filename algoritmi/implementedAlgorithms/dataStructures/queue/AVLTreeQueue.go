package queue

import (
	util "algoritmi/algoritmi/implementedAlgorithms/dataStructures"
)

type AVLTreeQueue struct {
	list []*util.AVLTree
}

func CreateAVLTreeQueue() *AVLTreeQueue {
	queue := new(AVLTreeQueue)
	return queue
}

func (treeNode *AVLTreeQueue) IsEmpty() bool { // Costo tempo => costante O(1)
	return len(treeNode.list) == 0
}

func (treeNode *AVLTreeQueue) Enqueue(val *util.AVLTree) {
	/*
		Costo tempo:
			Caso migliore => costante O(1)
			Caso peggiore => O(n)
			Questi costi sono dati dalla append che per come è implementata ha questa complessità
	*/
	treeNode.list = append(treeNode.list, val)
}

func (treeNode *AVLTreeQueue) First() *util.AVLTree { // Costo tempo => costante O(1)
	if !treeNode.IsEmpty() {
		return treeNode.list[0]
	}
	return nil
}

func (treeNode *AVLTreeQueue) Dequeue() *util.AVLTree { // Costo tempo => costante O(1)
	if !treeNode.IsEmpty() {
		val := treeNode.list[0]
		treeNode.list = treeNode.list[1:]
		return val
	}
	return nil
}
