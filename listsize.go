package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	length int
	Head   *NodeL
	Tail   *NodeL
}

func ListPushFront(l *List, data interface{}) {
	node := &NodeL{Data: data, Next: l.Head}
	l.Head = node
	if l.Tail == nil {
		l.Tail = l.Head
	}
	l.length++
}

func ListSize(l *List) int {
	return l.length
}
