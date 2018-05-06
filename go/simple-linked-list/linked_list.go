package linkedlist

import "errors"

// API:
//
// type Element struct {
//  data int
//  next *Element
// }
//
// type List struct {
//  head *Element
//  size int
// }
//
// func New([]int) *List
// func (*List) Size() int
// func (*List) Push(int)
// func (*List) Pop() (int, error)
// func (*List) Array() []int
// func (*List) Reverse() *List

type Element struct {
	data int
	next *Element
}

type List struct {
	head *Element
	size int
}

func New(a []int) *List {

	l := &List{head: nil, size: 0}
	for i := 0; i < len(a); i++ {
		l.Push(a[i])
	}
	l.size = len(a)
	return l
}

func (l *List) Size() int {
	return l.size
}

func (l *List) Push(i int) {
	l.size++
	if l.head == nil {
		l.head = &Element{data: i}
		return
	}
	for e := l.head; e != nil; e = e.next {
		if e.next == nil {
			e.next = &Element{data: i}
			return
		}
	}
}

func (l *List) Pop() (int, error) {
	if l.size == 0 {
		return 0, errors.New("empty list")
	}
	l.size--
	popped := &Element{}
	for e := l.head; e.next != nil; e = e.next {
		if e.next.next == nil {
			popped = e.next
			e.next = nil
			return popped.data, nil
		}
	}
	popped = l.head
	l.head = nil
	return popped.data, nil

}

func (l *List) Array() []int {
	array := make([]int, l.size)
	for e, i := l.head, 0; e != nil || i < len(array); e, i = e.next, i+1 {
		array[i] = e.data
	}
	return array
}

func (l *List) Reverse() *List {
	a := l.Array()
	nl := New(nil)
	for i := len(a) - 1; i >= 0; i-- {
		nl.Push(a[i])
	}
	return nl

}
