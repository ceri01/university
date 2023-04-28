package main

import (
	advancedOrder "algoritmi/algoritmi/algoritmi_implementati/ordinamento/tecniche_avanzate"
	baseOrder "algoritmi/algoritmi/algoritmi_implementati/ordinamento/tecniche_base"
	"algoritmi/algoritmi/algoritmi_implementati/struttureDati/linkedList"
	"algoritmi/algoritmi/algoritmi_implementati/struttureDati/orderedLinkedList"
	"fmt"
)

func main() {
	fmt.Println("##### Test algoritmi di ricerca e ordinamento #####")
	arrSelection := []int{3, 6, 1, 0, 22, 4, 3}
	arrInsertion := []int{25, 16, 2, 67, 25, 19, 6}
	arrBubble := []int{47, 12, 12, 7, 25, 1, 65}
	arrMerge := []int{16, 3, 45, 22, 252, 56, 5}
	arrQuick := []int{123, 431, 84, 732, 22, 1, -3, 32}

	baseOrder.Selection_sort(arrSelection)
	baseOrder.Insertion_sort(arrInsertion)
	baseOrder.Bubble_sort(arrBubble)
	advancedOrder.Merge_sort(arrMerge)
	advancedOrder.Quick_sort(arrQuick)

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
}
