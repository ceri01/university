package funzioni

import "strings"

func QuanteConA(parole []string) int {
	count := 0
	for _, w := range parole {
		if strings.Contains(w, "a") {
			count++
		}
	}
	return count
}
