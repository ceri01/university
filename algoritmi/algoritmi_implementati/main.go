package main

import (
	ordinamento2 "algoritmi/algoritmi/algoritmi_implementati/ordinamento/tecniche_avanzate"
	ordinamento "algoritmi/algoritmi/algoritmi_implementati/ordinamento/tecniche_base"
	"fmt"
)

func main() {
	arrSelection := []int{3, 6, 1, 0, 22, 4, 3}
	arrInsertion := []int{25, 16, 2, 67, 25, 19, 6}
	arrBubble := []int{47, 12, 12, 7, 25, 1, 65}
	arrMerge := []int{16, 3, 45, 22, 252, 56, 5}
	ordinamento.Selection_sort(arrSelection)
	ordinamento.Insertion_sort(arrInsertion)
	ordinamento.Bubble_sort(arrBubble)
	ordinamento2.Merge_sort(arrMerge)

	fmt.Println("Array ordinato con Selection sort")
	for _, el := range arrSelection {
		fmt.Printf("%d ", el)
	}
	fmt.Println("")
	fmt.Println("Array ordinato con Insertion sort")
	for _, el := range arrInsertion {
		fmt.Printf("%d ", el)
	}
	fmt.Println("")
	fmt.Println("Array ordinato con Bubble sort")
	for _, el := range arrBubble {
		fmt.Printf("%d ", el)
	}
	fmt.Println("")

	fmt.Println("Array ordinato con Merge sort")
	for _, el := range arrMerge {
		fmt.Printf("%d ", el)
	}
	fmt.Println("")
}
