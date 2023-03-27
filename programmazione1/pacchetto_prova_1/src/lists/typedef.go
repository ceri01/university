// Pacchetto per la gestione delle liste concatenate semplici di stringhe

package lists

// A src Node is made by a string (the node's Content) and a pointer to the next Node in the list (nil for the last node)
type Node struct {
	Content string
	Next *Node
}

// A list is identified with the pointer
type List = *Node