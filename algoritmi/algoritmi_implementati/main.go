package main

import (
	advancedOrder "algoritmi/algoritmi/algoritmi_implementati/ordinamento/tecniche_avanzate"
	baseOrder "algoritmi/algoritmi/algoritmi_implementati/ordinamento/tecniche_base"
	linkedList "algoritmi/algoritmi/algoritmi_implementati/struttureDati"
	"fmt"
)

func main() {
	arrSelection := []int{3, 6, 1, 0, 22, 4, 3}
	arrInsertion := []int{25, 16, 2, 67, 25, 19, 6}
	arrBubble := []int{47, 12, 12, 7, 25, 1, 65}
	arrMerge := []int{16, 3, 45, 22, 252, 56, 5}
	arrQuick := []int{123, 431, 84, 732, 22, 1, -3, 32}
	list := new(linkedList.LinkedList)

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

	linkedList.PrintList(list)
	linkedList.InsertHead(list, 10)
	linkedList.InsertHead(list, 5)
	linkedList.InsertHead(list, 3)
	linkedList.InsertTail(list, 4)
	linkedList.PrintList(list)

	isPresent, el := linkedList.SearchByKey(list, 5)
	fmt.Println(isPresent, el.Key)
	isPresent, el = linkedList.SearchByKey(list, 12)
	fmt.Println(isPresent, el) // se el è nil non si può chiamare .Key
	fmt.Println(linkedList.SearchByKey(nil, 12))

	fmt.Println(linkedList.SearchByPosition(list, 3))
	fmt.Println(linkedList.SearchByPosition(list, 4))
	fmt.Println(linkedList.SearchByPosition(list, -2))
	fmt.Println(linkedList.SearchByPosition(list, 1))
	fmt.Println(linkedList.SearchByPosition(nil, 1))

	/* test removeByVal (works)
	linkedList.PrintList(list)
	linkedList.RemoveByVal(list, 3)
	linkedList.PrintList(list)
	linkedList.RemoveByVal(list, 5)
	linkedList.PrintList(list)
	linkedList.RemoveByVal(list, 10)
	linkedList.PrintList(list)
	linkedList.RemoveByVal(list, 4)
	linkedList.PrintList(list)
	*/

	linkedList.PrintList(list)
	linkedList.RemoveByposition(list, -1)
	linkedList.PrintList(list)
	linkedList.RemoveByposition(list, -1)
	linkedList.PrintList(list)
	linkedList.RemoveByposition(list, -1)
	linkedList.PrintList(list)
	linkedList.RemoveByposition(list, -1)
	linkedList.PrintList(list)
	linkedList.RemoveByposition(list, -1)
	linkedList.PrintList(list)
}
