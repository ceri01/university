package gen

import (
	"math/rand"
)

// Generates a random lowercase latin letter ans returns it in the form of a 1-character string.
func genCharacter() string {
	return string('a' + rune(rand.Intn('z' - 'a' + 1)))
}

// Generates iteratively a palyndrome string of lowercase latin letters of length n
func IterativeGenPalyn(n int) string {
	t := ""
	if n % 2 != 0 {
		t = genCharacter()
	}
	for i := 0; i < n/2; i++ {
		c := genCharacter()
		t = c + t + c
	}
	return t
}
