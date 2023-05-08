package main

import (
	"fmt"
)

func kek(i int) int64 {
	if i == 1 || i == 2 {
		return 1
	}
	return kek(i-1) + kek(i-2)
}

func main() {
	fmt.Println(kek(7))
}
