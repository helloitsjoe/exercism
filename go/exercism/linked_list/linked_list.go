package linkedlist

import "fmt"

// Define List and Node types here.
// Note: The tests expect Node type to include an exported field with name Value to pass.
type List struct {
	head *Node
	tail *Node
}
type Node struct {
	Value interface{}
	next  *Node
	prev  *Node
}

func NewList(args ...interface{}) *List {
	l := &List{}
	for i, arg := range args {
		n := &Node{Value: arg}
		if i == 0 {
			l.head = n
			l.tail = n
		} else {
			prevTail := l.tail
			newTail := n
			prevTail.next = newTail
			newTail.prev = prevTail
			l.tail = newTail
		}
	}

	return l
}

func (n *Node) Next() *Node {
	// Need to check for nil?
	return n.next
}

func (n *Node) Prev() *Node {
	// Need to check for nil?
	return n.prev
}

func (l *List) Unshift(v interface{}) {
	newHead := &Node{Value: v}

	// Add to head
	if l.head == nil {
		l.head = newHead
		l.tail = newHead
		return
	}

	prevHead := l.head

	prevHead.prev = newHead
	newHead.next = prevHead
	l.head = newHead
}

func (l *List) Push(v interface{}) {
	// Add to tail
	newTail := &Node{Value: v}

	if l.tail == nil {
		l.tail = newTail
		l.head = newTail
		return
	}

	prevTail := l.tail

	prevTail.next = newTail
	newTail.prev = prevTail
	l.tail = newTail
}

func (l *List) Shift() (interface{}, error) {
	// Remove from head
	if l.head == nil {
		// TODO: Return error?
		return nil, nil
	}

	if l.head.next == nil {
		val := l.head.Value
		l.head = nil
		l.tail = nil
		return val, nil
	}

	prevHead := l.head

	newHead := prevHead.next
	newHead.prev = nil
	fmt.Println("newHead", newHead)
	l.head = newHead

	return prevHead.Value, nil
}

func (l *List) Pop() (interface{}, error) {
	if l.tail == nil {
		// TODO: Return error?
		return nil, nil
	}

	if l.tail.prev == nil {
		val := l.tail.Value
		l.head = nil
		l.tail = nil
		return val, nil
	}

	prevTail := l.tail
	newTail := prevTail.prev
	newTail.next = nil
	l.tail = newTail

	return prevTail.Value, nil
}

func (l *List) Reverse() {
	prevHead := l.head
	prevTail := l.tail

	swap(l.tail)

	l.head = prevTail
	l.tail = prevHead
}

func swap(node *Node) {
	if node == nil {
		return
	}

	tempPrev := node.prev
	tempNext := node.next

	if tempPrev != nil {
		node.prev = tempNext
		node.next = tempPrev

	} else {
		node.prev = tempNext
		node.next = nil
	}

	swap(node.next)
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Last() *Node {
	return l.tail
}
