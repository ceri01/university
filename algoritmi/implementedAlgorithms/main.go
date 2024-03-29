package main

import (
	util "algoritmi/algoritmi/implementedAlgorithms/dataStructures"
	"algoritmi/algoritmi/implementedAlgorithms/dataStructures/graph"
	"algoritmi/algoritmi/implementedAlgorithms/dataStructures/linkedList"
	"algoritmi/algoritmi/implementedAlgorithms/dataStructures/orderedLinkedList"
	"algoritmi/algoritmi/implementedAlgorithms/dataStructures/queue"
	"algoritmi/algoritmi/implementedAlgorithms/dataStructures/stack"
	tree "algoritmi/algoritmi/implementedAlgorithms/dataStructures/tree"
	_func "algoritmi/algoritmi/implementedAlgorithms/func"
	"algoritmi/algoritmi/implementedAlgorithms/sort"
	advancedOrder "algoritmi/algoritmi/implementedAlgorithms/sort/advanced"
	baseOrder "algoritmi/algoritmi/implementedAlgorithms/sort/base"
	"algoritmi/algoritmi/implementedAlgorithms/sort/withoutComparison"
	"fmt"
)

func main() {
	fmt.Println("##### Test algoritmi di search e sort #####")
	arrSelection := []int{3, 6, 1, 0, 22, 4, 3}
	arrInsertion := []int{25, 16, 2, 67, 25, 19, 6}
	arrBubble := []int{47, 12, 12, 7, 25, 1, 65}
	arrMerge := []int{16, 3, 45, 22, 252, 56, 5}
	arrQuick := []int{123, 431, 84, 732, 22, 1, -3, 32}

	baseOrder.SelectionSort(arrSelection)
	baseOrder.InsertionSort(arrInsertion)
	baseOrder.BubbleSort(arrBubble)
	advancedOrder.MergeSort(arrMerge)
	advancedOrder.QuickSort(arrQuick)

	fmt.Println("Array ordinato con Selectionsort")
	for _, el := range arrSelection {
		fmt.Printf("%d ", el)
	}
	fmt.Println("")
	fmt.Println("Array ordinato con Insertionsort")
	for _, el := range arrInsertion {
		fmt.Printf("%d ", el)
	}
	fmt.Println("")
	fmt.Println("Array ordinato con Bubblesort")
	for _, el := range arrBubble {
		fmt.Printf("%d ", el)
	}
	fmt.Println("")

	fmt.Println("Array ordinato con Mergesort")
	for _, el := range arrMerge {
		fmt.Printf("%d ", el)
	}
	fmt.Println("")

	fmt.Println("Array ordinato con Quicksort")
	for _, el := range arrQuick {
		fmt.Printf("%d ", el)
	}
	fmt.Println("")

	// Test linked linkedList
	fmt.Println("\n\n##### Test LinkedList #####")
	fmt.Println()
	lnkdList := new(linkedList.LinkedList)
	lnkdList1 := new(linkedList.LinkedList)

	fmt.Println("	Inserimento in testa e in coda")
	lnkdList.PrintList()
	lnkdList.InsertHead(10)
	lnkdList.InsertHead(5)
	lnkdList.InsertHead(3)
	lnkdList.InsertTail(4)
	lnkdList.PrintList()

	fmt.Println("	Ricerca per chiave")
	isPresent, el := lnkdList.SearchByKey(5)
	fmt.Println(isPresent, el.Key)
	isPresent, el = lnkdList.SearchByKey(12)
	fmt.Println(isPresent, el) // se el è nil non si può chiamare .Key
	fmt.Println(lnkdList.SearchByKey(12))

	fmt.Println("	Ricerca per posizione nella lista")
	fmt.Println(lnkdList.SearchByPosition(3))
	fmt.Println(lnkdList.SearchByPosition(4))
	fmt.Println(lnkdList.SearchByPosition(-2))
	fmt.Println(lnkdList.SearchByPosition(1))

	fmt.Println("	Rimozione per valore")
	lnkdList1.InsertHead(5)
	lnkdList1.InsertHead(3)
	lnkdList1.InsertHead(10)
	lnkdList1.InsertTail(4)
	lnkdList1.PrintList()
	lnkdList1.RemoveByVal(3)
	lnkdList1.PrintList()
	lnkdList1.RemoveByVal(5)
	lnkdList1.PrintList()
	lnkdList1.RemoveByVal(10)
	lnkdList1.PrintList()
	lnkdList1.RemoveByVal(4)
	lnkdList1.PrintList()

	fmt.Println("	Rimozione per posizione")
	lnkdList.PrintList()
	lnkdList.RemoveByPosition(0)
	lnkdList.PrintList()
	lnkdList.RemoveByPosition(1)
	lnkdList.PrintList()
	lnkdList.RemoveByPosition(12)
	lnkdList.PrintList()
	lnkdList.RemoveByPosition(-1)
	lnkdList.PrintList()
	lnkdList.RemoveByPosition(0)
	lnkdList.PrintList()

	// Test orderedLinkedList
	fmt.Println("\n\n##### Test OrderedLinkedList #####")
	fmt.Println()
	ordLnkdList := new(orderedLinkedList.OrderedLinkedList)
	ordLnkdList1 := new(orderedLinkedList.OrderedLinkedList)

	fmt.Println("	Inserimento")
	ordLnkdList.Insert(2)
	ordLnkdList.PrintList()
	ordLnkdList.Insert(6)
	ordLnkdList.PrintList()
	ordLnkdList.Insert(3)
	ordLnkdList.PrintList()
	ordLnkdList.Insert(-1)
	ordLnkdList.PrintList()
	ordLnkdList.Insert(5)
	ordLnkdList.PrintList()
	ordLnkdList.Insert(6)
	ordLnkdList.PrintList()

	fmt.Println("	Ricerca per chiave")
	ordIsPresent, el := ordLnkdList.SearchByKey(5)
	fmt.Println(ordIsPresent, el.Key)
	ordIsPresent, el = ordLnkdList.SearchByKey(12)
	fmt.Println(ordIsPresent, el) // se el è nil non si può chiamare .Key

	fmt.Println("	Ricerca per posizione nella lista")
	fmt.Println(ordLnkdList.SearchByPosition(3))
	fmt.Println(ordLnkdList.SearchByPosition(4))
	fmt.Println(ordLnkdList.SearchByPosition(-2))
	fmt.Println(ordLnkdList.SearchByPosition(20))
	fmt.Println(ordLnkdList.SearchByPosition(1))

	fmt.Println("	Rimozione per valore")
	ordLnkdList1.Insert(10)
	ordLnkdList1.Insert(5)
	ordLnkdList1.Insert(3)
	ordLnkdList1.Insert(4)
	ordLnkdList1.RemoveByVal(7)
	ordLnkdList1.PrintList()
	ordLnkdList1.PrintList()
	ordLnkdList1.RemoveByVal(3)
	ordLnkdList1.PrintList()
	ordLnkdList1.RemoveByVal(5)
	ordLnkdList1.PrintList()
	ordLnkdList1.RemoveByVal(10)
	ordLnkdList1.PrintList()
	ordLnkdList1.RemoveByVal(-2)
	ordLnkdList1.PrintList()
	ordLnkdList1.RemoveByVal(4)
	ordLnkdList1.PrintList()

	fmt.Println("	Rimozione per posizione")
	ordLnkdList.PrintList()
	ordLnkdList.RemoveByPosition(3)
	ordLnkdList.PrintList()
	ordLnkdList.RemoveByPosition(1)
	ordLnkdList.PrintList()
	ordLnkdList.RemoveByPosition(-2)
	ordLnkdList.PrintList()
	ordLnkdList.RemoveByPosition(22)
	ordLnkdList.PrintList()
	ordLnkdList.RemoveByPosition(1)
	ordLnkdList.PrintList()
	ordLnkdList.RemoveByPosition(2)
	ordLnkdList.PrintList()
	ordLnkdList.RemoveByPosition(0)
	ordLnkdList.PrintList()

	// Test array stack
	fmt.Println("\n\n##### Test Array Stack #####")
	fmt.Println()

	arrStk := stack.CreateArrayStack()
	fmt.Println(arrStk.IsEmpty())
	arrStk.PrintStack()
	arrStk.Push(2)
	arrStk.Push(5)
	// Se lo stack viene implementato con una dimensione troppo bassa (es: 3) e vengono aggiunti degli elementi si avrà un errore.
	// Ma è normale che sia cosi, è proprio un problema implementativo
	// arrStk.Enqueue(5)
	// arrStk.Enqueue(5)
	arrStk.PrintStack()
	fmt.Println(arrStk.IsEmpty())
	fmt.Println(arrStk.Top())
	arrStk.PrintStack()
	arrStk.Pop()
	arrStk.PrintStack()
	fmt.Println(arrStk.Top())

	// Test list stack
	fmt.Println("\n\n##### Test List Stack #####")
	fmt.Println()

	lstStk := stack.CreateListStack()
	fmt.Println(lstStk.IsEmpty())
	lstStk.PrintStack()
	lstStk.Push(2)
	lstStk.Push(5)
	lstStk.PrintStack()
	fmt.Println(lstStk.IsEmpty())
	fmt.Println(lstStk.Top())
	lstStk.PrintStack()
	lstStk.Pop()
	lstStk.PrintStack()
	fmt.Println(lstStk.Top())

	// Test Customlist stack
	fmt.Println("\n\n##### Test Custom List Stack #####")
	fmt.Println()
	cstLstStk := stack.CreateCustomListStack()
	fmt.Println(cstLstStk.IsEmpty())
	cstLstStk.PrintStack()
	cstLstStk.Push(2)
	cstLstStk.Push(5)
	cstLstStk.Push(2)
	cstLstStk.Push(4)
	cstLstStk.Push(2)
	cstLstStk.Push(9)
	cstLstStk.PrintStack()
	fmt.Println(cstLstStk.IsEmpty())
	fmt.Println(cstLstStk.Top())
	cstLstStk.PrintStack()
	cstLstStk.Pop()
	cstLstStk.PrintStack()
	fmt.Println(cstLstStk.Top())

	// Test linked doubleLinkedList
	fmt.Println("\n\n##### Test DoubleLinkedList #####")
	fmt.Println()
	dlnkdList := new(linkedList.DoubleLinkedList)
	dlnkdList1 := new(linkedList.DoubleLinkedList)

	fmt.Println("	Inserimento in testa e in coda")
	dlnkdList.PrintList()
	dlnkdList.InsertHead(10)
	dlnkdList.InsertTail(1)
	dlnkdList.InsertHead(5)
	dlnkdList.InsertHead(3)
	dlnkdList.InsertTail(4)
	dlnkdList.PrintList()

	fmt.Println("	Ricerca per chiave")
	isPresent, eld := dlnkdList.SearchByKey(5)
	fmt.Println(isPresent, eld.Key)
	isPresent, eld = dlnkdList.SearchByKey(12)
	fmt.Println(isPresent, eld) // se el è nil non si può chiamare .Key
	fmt.Println(dlnkdList.SearchByKey(12))

	fmt.Println("	Ricerca per posizione nella lista")
	dlnkdList.PrintList()
	fmt.Println(dlnkdList.SearchByPosition(3))
	fmt.Println(dlnkdList.SearchByPosition(5))
	fmt.Println(dlnkdList.SearchByPosition(-2))
	fmt.Println(dlnkdList.SearchByPosition(1))

	fmt.Println("	Rimozione per valore")
	dlnkdList1.InsertHead(5)
	dlnkdList1.InsertHead(3)
	dlnkdList1.InsertHead(10)
	dlnkdList1.InsertTail(4)
	dlnkdList1.PrintList()
	dlnkdList1.RemoveByVal(3)
	dlnkdList1.PrintList()
	dlnkdList1.RemoveByVal(5)
	dlnkdList1.PrintList()
	dlnkdList1.RemoveByVal(10)
	dlnkdList1.PrintList()
	dlnkdList1.RemoveByVal(4)
	dlnkdList1.PrintList()

	fmt.Println("	Rimozione per posizione")
	dlnkdList.PrintList()
	dlnkdList.RemoveByPosition(0)
	dlnkdList.PrintList()
	dlnkdList.RemoveByPosition(1)
	dlnkdList.PrintList()
	dlnkdList.RemoveByPosition(12)
	dlnkdList.PrintList()
	dlnkdList.RemoveByPosition(-1)
	dlnkdList.PrintList()
	dlnkdList.RemoveByPosition(0)
	dlnkdList.PrintList()

	// Test array stack
	fmt.Println("\n\n##### Test Array Queue #####")
	fmt.Println()

	arrQue := queue.CreateArrayQueue()
	fmt.Println(arrQue.IsEmpty())
	arrQue.PrintQueue()
	arrQue.Enqueue(2)
	arrQue.Enqueue(5)
	arrQue.Enqueue(1)
	arrQue.Enqueue(7)
	arrQue.Enqueue(4)
	arrQue.PrintQueue()
	arrQue.Dequeue()
	arrQue.PrintQueue()
	arrQue.Dequeue()
	arrQue.PrintQueue()
	arrQue.Dequeue()
	arrQue.PrintQueue()
	arrQue.Dequeue()
	arrQue.PrintQueue()
	arrQue.Dequeue()
	arrQue.PrintQueue()
	arrQue.Dequeue()
	arrQue.PrintQueue()
	arrQue.Enqueue(7)
	arrQue.Enqueue(4)
	arrQue.PrintQueue()
	arrQue.Dequeue()
	arrQue.PrintQueue()
	fmt.Println(arrQue.First())

	// Test List stack
	fmt.Println("\n\n##### Test List Queue #####")
	fmt.Println()

	lstQue := queue.CreateListQueue()
	fmt.Println(lstQue.IsEmpty())
	lstQue.PrintQueue()
	lstQue.Enqueue(2)
	lstQue.Enqueue(5)
	lstQue.Enqueue(1)
	lstQue.Enqueue(7)
	lstQue.Enqueue(4)
	lstQue.PrintQueue()
	lstQue.Dequeue()
	lstQue.PrintQueue()
	lstQue.Dequeue()
	lstQue.PrintQueue()
	lstQue.Dequeue()
	lstQue.PrintQueue()
	lstQue.Dequeue()
	lstQue.PrintQueue()
	lstQue.Dequeue()
	lstQue.PrintQueue()
	lstQue.Dequeue()
	lstQue.PrintQueue()
	fmt.Println(lstQue.First())
	lstQue.Enqueue(7)
	lstQue.Enqueue(4)
	lstQue.PrintQueue()
	lstQue.Dequeue()
	lstQue.PrintQueue()
	fmt.Println(lstQue.First())

	// Test Custom List stack
	fmt.Println("\n\n##### Test Custom List Queue #####")
	fmt.Println()

	cstLstQue := queue.CreateListQueue()
	fmt.Println(cstLstQue.IsEmpty())
	cstLstQue.PrintQueue()
	cstLstQue.Enqueue(2)
	cstLstQue.Enqueue(5)
	cstLstQue.Enqueue(1)
	cstLstQue.Enqueue(7)
	cstLstQue.Enqueue(4)
	cstLstQue.PrintQueue()
	cstLstQue.Dequeue()
	cstLstQue.PrintQueue()
	cstLstQue.Dequeue()
	cstLstQue.PrintQueue()
	cstLstQue.Dequeue()
	cstLstQue.PrintQueue()
	cstLstQue.Dequeue()
	cstLstQue.PrintQueue()
	cstLstQue.Dequeue()
	cstLstQue.PrintQueue()
	cstLstQue.Dequeue()
	cstLstQue.PrintQueue()
	fmt.Println(cstLstQue.First())
	cstLstQue.Enqueue(7)
	cstLstQue.Enqueue(4)
	cstLstQue.PrintQueue()
	cstLstQue.Dequeue()
	cstLstQue.PrintQueue()
	fmt.Println(cstLstQue.First())

	// Test Binary Tree view
	fmt.Println("\n\n##### Test Binary tree view #####")

	t := tree.CreateBinaryTree()
	t.Root = &util.BiTreeNode{Val: 50, Left: nil, Right: nil}
	t.Root.Left = util.NewTreeNode(20)
	t.Root.Right = util.NewTreeNode(80)
	t.Root.Left.Left = util.NewTreeNode(10)
	t.Root.Left.Right = util.NewTreeNode(8)
	t.Root.Right.Left = util.NewTreeNode(45)
	t.Root.Right.Right = util.NewTreeNode(2)
	t.Root.Left.Left.Left = util.NewTreeNode(4)
	t.Root.Left.Left.Right = util.NewTreeNode(43)
	t.Root.Left.Right.Left = util.NewTreeNode(67)
	t.Root.Left.Right.Right = util.NewTreeNode(22)
	t.Root.Right.Left.Left = util.NewTreeNode(453)
	t.Root.Right.Left.Right = util.NewTreeNode(32)
	t.Root.Right.Right.Left = util.NewTreeNode(88)
	t.Root.Right.Right.Right = util.NewTreeNode(7)

	t.BFS()
	t.DFS(tree.InOrder)
	t.DFS(tree.PreOrder)
	t.DFS(tree.PostOrder)
	fmt.Println(t.Nnode(t.Root))

	// Test Binary Search Tree
	fmt.Println("\n\n##### Test Binary Search tree #####")
	st := tree.CreateBinarySearchTree(8)
	tree.InsertRecursive(st, 4)
	tree.InsertRecursive(st, 0)
	tree.InsertRecursive(st, -1)
	tree.InsertRecursive(st, 6)
	tree.InsertRecursive(st, 7)
	tree.InsertIterative(st, 15)
	tree.InsertIterative(st, 10)
	tree.InsertIterative(st, 11)
	tree.InsertIterative(st, 19)
	tree.InsertIterative(st, 20)

	tree.DFS(st, tree.InOrderSearchBiTree)
	tree.BFS(st)
	fmt.Println(tree.Max(st))
	tree.SummaryView(st, 0)
	fmt.Println(tree.FindRecursive(st, 45))
	fmt.Println(tree.FindRecursive(st, 90))
	fmt.Println(tree.FindIterative(st, -1))
	fmt.Println(tree.FindIterative(st, 50))

	fmt.Println("##DELETES BINARY SEARCH TREE##")
	// delete node with 1 child
	fmt.Println("cancel 7")
	tree.Remove(st, 7)
	tree.SummaryView(st, 0)
	tree.InsertRecursive(st, 7)
	fmt.Println("cancel 0")
	tree.Remove(st, 0)
	tree.SummaryView(st, 0)
	tree.InsertRecursive(st, 0)
	fmt.Println("cancel 10")
	tree.Remove(st, 10)
	tree.SummaryView(st, 0)
	tree.InsertRecursive(st, 10)
	fmt.Println("cancel 19")
	tree.Remove(st, 19)
	tree.SummaryView(st, 0)

	// delete leaf
	fmt.Println("cancel 20")
	tree.Remove(st, 20)
	tree.SummaryView(st, 0)

	// delete node with 2 children
	fmt.Println("cancel 4")
	tree.Remove(st, 4)
	tree.SummaryView(st, 0)

	// delete root
	fmt.Println("cancel 8")
	tree.Remove(st, 8)
	tree.SummaryView(st, 0)

	// Test AVL-Tree
	fmt.Println("\n\n##### Test AVL-Tree #####")
	avl := tree.CreateAVLTree(30)
	tree.Insert(avl, 50)
	tree.Insert(avl, 10)
	tree.SummaryViewAVL(avl, 0)
	tree.Insert(avl, 15)
	tree.Insert(avl, 6)
	tree.SummaryViewAVL(avl, 0)
	tree.Insert(avl, 20)
	tree.SummaryViewAVL(avl, 0)

	// Test Heap
	fmt.Println("\n\n##### Test Heap and heapsort #####")
	heap := []int{4, 5, 6, 7, 8, 9, 10}
	tree.HeapSort(heap)
	for i := 0; i < len(heap); i++ {
		fmt.Printf("%d ", heap[i])
	}
	fmt.Println()

	heap = []int{43, 123, 994, 5, 5, 5, 5, 123, 6, 22, 5, 5, 22, 5}
	tree.HeapSort(heap)
	for i := 0; i < len(heap); i++ {
		fmt.Printf("%d ", heap[i])
	}
	fmt.Println()

	// Test IntegerSort
	fmt.Println("\n\n##### Test IntegerSort #####")
	insertionData := []int{9, 10, 5, 1, 6, 8, 0, 4, 0}
	// 11 è il valore dell'elemento più grande aumentato di uno (quindi gl elementi vanno da zero a undici)
	withoutComparison.IntegerSort(insertionData, insertionData[_func.RicercaMax(insertionData)]+1)
	fmt.Println(insertionData)

	// Test BucketSort
	fmt.Println("\n\n##### Test BucketSort #####")
	bucketData := []*sort.Record{
		{Key: 2, Data: "Cane"},
		{Key: 1, Data: "Gatto"},
		{Key: 5, Data: "Mucca"},
		{Key: 2, Data: "Gallina"},
		{Key: 7, Data: "Toro"},
		{Key: 5, Data: "Zanzara"},
		{Key: 0, Data: "Mulo"},
	}
	withoutComparison.BucketSort(bucketData, 7)
	limit := len(bucketData)
	for i := 0; i < limit; i++ {
		fmt.Printf("%d => %s\n", bucketData[i].Key, bucketData[i].Data)
	}

	// Test RadixData !!! NON FUNZIONA !!!
	/*	fmt.Println("\n\n##### Test RadixData #####")
		radixData := []*sort.Record{
			{Key: 45, Data: "Cane"},
			{Key: 23, Data: "Gatto"},
			{Key: 123, Data: "Mucca"},
			{Key: 44, Data: "Gallina"},
			{Key: 5, Data: "Toro"},
			{Key: 11, Data: "Zanzara"},
			{Key: 5, Data: "Mulo"},
		}
		withoutComparison.RadixSort(radixData)
		limit = len(radixData)
		for i := 0; i < limit; i++ {
			fmt.Printf("%d => %s\n", radixData[i].Key, radixData[i].Data)
		}*/

	// Test CompolexAdjacencyGraphs
	fmt.Println("\n\n##### Test complex adjacency graphs #####")
	fmt.Println("Uncomment to test it (works)")
	/*	cag := graph.NewComplexGraph(5)
		graph.ReadComplexGraph(cag)
		graph.PrintComplexGraph(cag)*/

	// Test MapAdjacencyListGraphs
	fmt.Println("\n\n##### Test Map Adjacency List Graphs #####")
	fmt.Println("Uncomment to test it (works)")
	/*	malg := graph.NewMapAdjListGraph(5)
		graph.ReadMapAdjListGraph(malg)
		graph.PrintMapAdjListGraph(malg)*/

	// Test AdjacencyListGraphs
	fmt.Println("\n\n##### Test Adjacency List Graphs #####")
	fmt.Println("Uncomment to test functions until BFS (works)")
	/*	alg := graph.NewAdjListGraph(10)
		graph.ReadAdjListGraph(alg)
		graph.PrintAdjListGraph(alg)
		resTree := graph.DFS(alg, 0)
		resTree = graph.BFS(alg, 0)
		tree.SummaryView(resTree, 0)
	*/
	alg1 := graph.NewAdjListGraph(5)
	graph.Gen(alg1, 0.2)
	graph.PrintAdjListGraph(alg1)
	fmt.Println("deg of 0: ", graph.Degree(alg1, 0))
	fmt.Println("deg of 6: ", graph.Degree(alg1, 6))
	fmt.Println("path from 0 to 3: ", graph.Path(alg1, 0, 3))
	// Test adjacencyMatrix
	fmt.Println("\n\n##### Test Adjacency matrix Graphs #####")
	fmt.Println("Uncomment to test it (works)")
	/*	amg := graph.NewAdjMatrixGraph(5)
		graph.ReadAdjMatrixGraph(amg)
		graph.PrintAdjMatrixGraph(amg)*/

	// Test adjacencyMatrixNotOriented
	fmt.Println("\n\n##### Test Adjacency List Graphs not oriented #####")
	/*	amgno := graph.NewAdjMatrixGraphNotOriented(5)
		graph.ReadAdjMatrixGraphNotOriented(amgno)
		graph.PrintAdjMatrixGraphNotOriented(amgno)*/
}
