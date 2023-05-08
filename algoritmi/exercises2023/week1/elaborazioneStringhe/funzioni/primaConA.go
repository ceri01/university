package funzioni

import "strings"

func PrimaConA(parole []string) string {
	for _, w := range parole {
		if strings.Contains(w, "a") {
			return w
		}
	}
	return ""
}
