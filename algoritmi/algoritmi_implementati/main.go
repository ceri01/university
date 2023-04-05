package main

import (
	ordinamento "algoritmi/algoritmi/algoritmi_implementati/ordinamento/tecniche_base"
	"fmt"
)

func main() {
	arr := []int{3, 6, 1, 0, 22, 4, 3}
	ordinamento.Selection_sort(arr)
	for _, el := range arr {
		fmt.Println(el)
	}

}
