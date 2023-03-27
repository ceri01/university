// The playn/check package contains functions that check if a given string is palyndrome or not. The
//check can be perfdormed iteratively or recursive

package check

// Checks (iteratively) if the string s is palyndrome, and returns the corresponding boolean value. The string
// is assumed to be an ASCII strings; the behavior on non-ASCII strings in unspacified

func RecIsPalyn(s string) bool {
	if len(s) < 2 {
		return true
	}
	return s[0] == s[len(s) - 1] && RecIsPalyn(s[1:len(s) - 1])
}
