package main

import "fmt"

type List[T any] struct {
	next *List[T]
	val  T
}

func (l *List[T]) AddLast(val T) *List[T] {
	if l == nil {
		l = &List[T]{val: val}
		return l
	}

	curr := l
	for curr.next != nil {
		curr = curr.next
	}

	curr.next = &List[T]{val: val}
	return l
}

func (l *List[T]) Iterate() {
	for cur := l; cur != nil; cur = cur.next {
		fmt.Println(cur.val)
	}
}

func main() {
	var l *List[int]
	l = l.AddLast(10)
	l = l.AddLast(20)
	l.Iterate()
}
