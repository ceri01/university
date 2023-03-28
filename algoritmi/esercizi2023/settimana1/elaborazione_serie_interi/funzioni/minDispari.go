package funzioni

import (
	"math"
)

func MinDispari(numeri []int) int {
	min := math.MaxInt64
	for _, val := range numeri {
		if val%2 != 0 && val < min {
			min = val
		}
	}
	return min
}
