package queue

import (
	util "algoritmi/algoritmi/implementedAlgorithms/sort"
)

type RecordQueue struct {
	list []*util.Record
}

func CreateRecordQueue() *RecordQueue {
	queue := new(RecordQueue)
	return queue
}

func (treeNode *RecordQueue) IsEmpty() bool { // Costo tempo => costante O(1)
	return len(treeNode.list) == 0
}

func (treeNode *RecordQueue) Enqueue(val *util.Record) {
	/*
		Costo tempo:
			Caso migliore => costante O(1)
			Caso peggiore => O(n)
			Questi costi sono dati dalla append che per come è implementata ha questa complessità
	*/
	treeNode.list = append(treeNode.list, val)
}

func (treeNode *RecordQueue) First() *util.Record { // Costo tempo => costante O(1)
	if !treeNode.IsEmpty() {
		return treeNode.list[0]
	}
	return nil
}

func (treeNode *RecordQueue) Dequeue() *util.Record { // Costo tempo => costante O(1)
	if !treeNode.IsEmpty() {
		val := treeNode.list[0]
		treeNode.list = treeNode.list[1:]
		return val
	}
	return nil
}
