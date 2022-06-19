package linkedlist

import (
	"errors"
)

type List struct {
	head *Element
	tail *Element
	size int
}

type Element struct {
	value int
	next  *Element
}

func New(input []int) *List {
	l := &List{}

	for _, value := range input {
		l.Push(value)
	}

	return l
}

func (l *List) Size() int {
	return l.size
}

func (l *List) Push(element int) {
	e := &Element{value: element}

	if l.head == nil {
		l.head = e
	} else {
		l.tail.next = e
	}

	l.tail = e
	l.size += 1
}

func (l *List) Pop() (int, error) {
	if l.head == nil {
		return 0, errors.New("List is empty, could not pop")
	}

	prev := l.head
	curr := l.head

	for curr.next != nil {
		prev = curr
		curr = curr.next
	}

	l.tail = prev
	l.tail.next = nil
	l.size -= 1
	return curr.value, nil
}

func (l *List) Array() []int {
	output := make([]int, l.size)

	for i := l.size - 1; i >= 0; i -= 1 {
		output[i], _ = l.Pop()
	}

	return output
}

func (l *List) Reverse() *List {
	a := l.Array()
	r := make([]int, len(a))

	for i := range a {
		r[i] = a[len(a)-1-i]
	}

	return New(r)
}
