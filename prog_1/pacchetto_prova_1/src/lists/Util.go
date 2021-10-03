package lists

// Returns the lenght (number if nodes) of x
func (x List) Length() (c int) {
	for x != nil {
		x = x.Next
		c++
	}
	return c
}
