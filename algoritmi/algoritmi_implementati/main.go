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
	linkedList.PrintList(lnkdList)
	linkedList.InsertHead(lnkdList, 10)
	linkedList.InsertHead(lnkdList, 5)
	linkedList.InsertHead(lnkdList, 3)
	linkedList.InsertTail(lnkdList, 4)
	linkedList.PrintList(lnkdList)

	fmt.Println("	Ricerca per chiave")
	isPresent, el := linkedList.SearchByKey(lnkdList, 5)
	fmt.Println(isPresent, el.Key)
	isPresent, el = linkedList.SearchByKey(lnkdList, 12)
	fmt.Println(isPresent, el) // se el è nil non si può chiamare .Key
	fmt.Println(linkedList.SearchByKey(nil, 12))

	fmt.Println("	Ricerca per posizione nella lista")
	fmt.Println(linkedList.SearchByPosition(lnkdList, 3))
	fmt.Println(linkedList.SearchByPosition(lnkdList, 4))
	fmt.Println(linkedList.SearchByPosition(lnkdList, -2))
	fmt.Println(linkedList.SearchByPosition(lnkdList, 1))
	fmt.Println(linkedList.SearchByPosition(nil, 1))

	fmt.Println("	Rimozione per valore")
	linkedList.InsertHead(lnkdList1, 10)
	linkedList.InsertHead(lnkdList1, 5)
	linkedList.InsertHead(lnkdList1, 3)
	linkedList.InsertTail(lnkdList1, 4)
	linkedList.PrintList(lnkdList1)
	linkedList.RemoveByVal(lnkdList1, 3)
	linkedList.PrintList(lnkdList1)
	linkedList.RemoveByVal(lnkdList1, 5)
	linkedList.PrintList(lnkdList1)
	linkedList.RemoveByVal(lnkdList1, 10)
	linkedList.PrintList(lnkdList1)
	linkedList.RemoveByVal(lnkdList1, 4)
	linkedList.PrintList(lnkdList1)

	fmt.Println("	Rimozione per posizione")
	linkedList.PrintList(lnkdList)
	linkedList.RemoveByPosition(lnkdList, 0)
	linkedList.PrintList(lnkdList)
	linkedList.RemoveByPosition(lnkdList, 1)
	linkedList.PrintList(lnkdList)
	linkedList.RemoveByPosition(lnkdList, 12)
	linkedList.PrintList(lnkdList)
	linkedList.RemoveByPosition(lnkdList, -1)
	linkedList.PrintList(lnkdList)
	linkedList.RemoveByPosition(lnkdList, 0)
	linkedList.PrintList(lnkdList)

	// Test orderedLinkedList
	fmt.Println("\n\n##### Test OrderedLinkedList #####")
	fmt.Println()
	ordLnkdList := new(orderedLinkedList.OrderedLinkedList)
	ordLnkdList1 := new(orderedLinkedList.OrderedLinkedList)

	fmt.Println("	Inserimento")
	orderedLinkedList.Insert(ordLnkdList, 2)
	orderedLinkedList.PrintList(ordLnkdList)
	orderedLinkedList.Insert(ordLnkdList, 6)
	orderedLinkedList.PrintList(ordLnkdList)
	orderedLinkedList.Insert(ordLnkdList, 3)
	orderedLinkedList.PrintList(ordLnkdList)
	orderedLinkedList.Insert(ordLnkdList, -1)
	orderedLinkedList.PrintList(ordLnkdList)
	orderedLinkedList.Insert(ordLnkdList, 5)
	orderedLinkedList.PrintList(ordLnkdList)
	orderedLinkedList.Insert(ordLnkdList, 6)
	orderedLinkedList.PrintList(ordLnkdList)

	fmt.Println("	Ricerca per chiave")
	ordIsPresent, el := orderedLinkedList.SearchByKey(ordLnkdList, 5)
	fmt.Println(ordIsPresent, el.Key)
	ordIsPresent, el = orderedLinkedList.SearchByKey(ordLnkdList, 12)
	fmt.Println(ordIsPresent, el) // se el è nil non si può chiamare .Key
	fmt.Println(orderedLinkedList.SearchByKey(nil, 12))

	fmt.Println("	Ricerca per posizione nella lista")
	fmt.Println(orderedLinkedList.SearchByPosition(ordLnkdList, 3))
	fmt.Println(orderedLinkedList.SearchByPosition(ordLnkdList, 4))
	fmt.Println(orderedLinkedList.SearchByPosition(ordLnkdList, -2))
	fmt.Println(orderedLinkedList.SearchByPosition(ordLnkdList, 20))
	fmt.Println(orderedLinkedList.SearchByPosition(nil, 1))

	fmt.Println("	Rimozione per valore")
	orderedLinkedList.Insert(ordLnkdList1, 10)
	orderedLinkedList.Insert(ordLnkdList1, 5)
	orderedLinkedList.Insert(ordLnkdList1, 3)
	orderedLinkedList.Insert(ordLnkdList1, 4)
	orderedLinkedList.RemoveByVal(ordLnkdList1, 7)
	orderedLinkedList.PrintList(ordLnkdList1)
	orderedLinkedList.PrintList(ordLnkdList1)
	orderedLinkedList.RemoveByVal(ordLnkdList1, 3)
	orderedLinkedList.PrintList(ordLnkdList1)
	orderedLinkedList.RemoveByVal(ordLnkdList1, 5)
	orderedLinkedList.PrintList(ordLnkdList1)
	orderedLinkedList.RemoveByVal(ordLnkdList1, 10)
	orderedLinkedList.PrintList(ordLnkdList1)
	orderedLinkedList.RemoveByVal(ordLnkdList1, -2)
	orderedLinkedList.PrintList(ordLnkdList1)

	fmt.Println("	Rimozione per posizione")
	orderedLinkedList.PrintList(ordLnkdList)
	orderedLinkedList.RemoveByPosition(ordLnkdList, 3)
	orderedLinkedList.PrintList(ordLnkdList)
	orderedLinkedList.RemoveByPosition(ordLnkdList, 1)
	orderedLinkedList.PrintList(ordLnkdList)
	orderedLinkedList.RemoveByPosition(ordLnkdList, -2)
	orderedLinkedList.PrintList(ordLnkdList)
	orderedLinkedList.RemoveByPosition(ordLnkdList, 22)
	orderedLinkedList.PrintList(ordLnkdList)
	orderedLinkedList.RemoveByPosition(ordLnkdList, 1)
	orderedLinkedList.PrintList(ordLnkdList)
	orderedLinkedList.RemoveByPosition(ordLnkdList, 2)
	orderedLinkedList.PrintList(ordLnkdList)
	orderedLinkedList.RemoveByPosition(ordLnkdList, 0)
	orderedLinkedList.PrintList(ordLnkdList)
}
