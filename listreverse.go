package piscine

func ListReverse(l *List) {
	list := &List{}
	for l.Head != nil {
		ListPushBack(list, l.Head.Data)
		l.Head = l.Head.Next
	}
	l.Head, l.Tail = list.Head, list.Tail
}
