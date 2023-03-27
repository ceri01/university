// The playn/check package contains functions that check if a given string is palyndrome or not. The
//check can be perfdormed iteratively or recursive

package check

// Checks (iteratively) if the string s is palyndrome, and returns the corresponding boolean value. The string
// is assumed to be an ASCII strings; the behavior on non-ASCII strings in unspacified

func IterIsPalyn(s string) bool {
	n := len(s)
	for i := 0; i < n / 2; i++ {
		if s[i] != s[n - i - 1] {
			return false
		}
	}
	return true
}