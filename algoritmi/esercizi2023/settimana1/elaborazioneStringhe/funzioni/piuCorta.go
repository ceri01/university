package funzioni

import "math"

func PiuCorta(parole []string) int {
	var shortest = math.MaxInt
	for _, w := range parole {
		if len(w) < shortest {
			shortest = len(w)
		}
	}
	return shortest
}
