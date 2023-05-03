package struttureDati

type ListNode struct {
	Key  int // In questo caso Key è la chiave, ma possiamo aggiungere altri campi
	Next *ListNode
}

type ListNodeDouble struct {
	Key  int // In questo caso Key è la chiave, ma possiamo aggiungere altri campi
	Next *ListNodeDouble
	Prev *ListNodeDouble
}

func NewNode(val int) *ListNode {
	// non è necessario deferenziare i puntatori in Go, infatti non è necessario dichiarare var node = *listNode
	// e utilizzarlo come in C.
	node := new(ListNode)
	node.Key = val
	node.Next = nil
	return node
	// node = &listNode{val, nil} => versione più breve delle tre righe sopra
}

func NewDoubleNode(val int) *ListNodeDouble {
	// non è necessario deferenziare i puntatori in Go, infatti non è necessario dichiarare var node = *listNode
	// e utilizzarlo come in C.
	node := new(ListNodeDouble)
	node.Key = val
	node.Next = nil
	node.Prev = nil
	return node
	// node = &listNode{val, nil} => versione più breve delle tre righe sopra
}