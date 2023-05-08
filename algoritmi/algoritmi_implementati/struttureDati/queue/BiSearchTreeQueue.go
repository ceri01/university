package queue

import (
	util "algoritmi/algoritmi/algoritmi_implementati/struttureDati"
)

type BiSearchTreeQueue struct {
	list []*util.BiSearchTree
}

func CreateBiSearchTreeQueue() *BiSearchTreeQueue {
	queue := new(BiSearchTreeQueue)
	return queue
}

func (treeNode *BiSearchTreeQueue) IsEmpty() bool { // Costo tempo => costante O(1)
	return len(treeNode.list) == 0
}

func (treeNode *BiSearchTreeQueue) Enqueue(val *util.BiSearchTree) { // Costo tempo => costante O(1)
	treeNode.list = append(treeNode.list, val)
}

func (treeNode *BiSearchTreeQueue) First() *util.BiSearchTree { // Costo tempo => costante O(1)
	if !treeNode.IsEmpty() {
		return treeNode.list[0]
	}
	return nil
}

func (treeNode *BiSearchTreeQueue) Dequeue() *util.BiSearchTree { // Costo tempo => costante O(1)
	if !treeNode.IsEmpty() {
		val := treeNode.list[0]
		treeNode.list = treeNode.list[1:]
		return val
	}
	return nil
}
