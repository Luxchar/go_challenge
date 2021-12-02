package main

import (
	"fmt"
)

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	node := &NodeL{Data: data, Next: nil}
	if l.Tail == nil {
		l.Head = node
		l.Tail = node
	} else {
		l.Tail.Next = node
		l.Tail = node
	}
}

func ListReverse(l *List) {
	cur := l.Head
	pre := l.Head
	pre = nil
	for cur != nil {
		pre, cur, cur.Next = cur, cur.Next, pre
	}
	l.Head, l.Tail = l.Tail, l.Head
}

func main() {
	link := &List{}

	ListPushBack(link, 1)
	ListPushBack(link, 2)

	ListPushBack(link, 3)
	ListPushBack(link, 4)

	ListReverse(link)

	it := link.Head

	for it != nil {
		fmt.Println(it.Data)
		it = it.Next
	}

	fmt.Println("Tail", link.Tail)
	fmt.Println("Head", link.Head)
}
