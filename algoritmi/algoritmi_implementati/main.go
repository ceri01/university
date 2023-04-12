package main

import (
	ordinamento "algoritmi/algoritmi/algoritmi_implementati/ordinamento/tecniche_base"
	"fmt"
)

func main() {
	arrSelection := []int{3, 6, 1, 0, 22, 4, 3}
	arrInsertion := []int{25, 16, 2, 67, 25, 19, 6}
	ordinamento.Selection_sort(arrSelection)
	ordinamento.Insertion_sort(arrInsertion)

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
}
