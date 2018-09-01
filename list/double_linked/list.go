package double_linked

import (
	"fmt"
	"strings"
)

// List is a container structer for list items
type List struct {
	first *Node
	last  *Node

	len int
}

// Node is an implementation of list item
type Node struct {
	value interface{}
	next  *Node
	prev  *Node
}

// Next returns next item or nil if no exists
func (n *Node) Next() *Node {
	return n.next
}

// Prev returns prev item or nil if no exists
func (n *Node) Prev() *Node {
	return n.prev
}

// NewList returns new LinkedList instance
func NewList() *List {
	ll := &List{}
	return ll
}

// Reverse reverses list items
// Reverse is mutating operation
func (l *List) Reverse() *List {
	n := l.first

	var prev *Node
	var next *Node

	for n != nil {
		next = n.next
		n.next = prev
		n.prev = next

		prev = n
		n = next
	}

	h := l.first
	t := l.last

	l.first = t
	l.last = h

	return l
}

// PushBack pushes new item to list end
func (l *List) PushBack(v interface{}) *Node {
	l.len = l.len + 1

	if l.first == nil {
		l.first = &Node{value: v}
		l.last = l.first

		return l.first
	}

	node := &Node{value: v}
	l.last.next = node

	node.prev = l.last

	l.last = node

	return node
}

// PushFront add new item in front of all items
func (l *List) PushFront(v interface{}) *Node {
	l.len = l.len + 1

	if l.first == nil {
		l.first = &Node{value: v}
		l.last = l.first

		return l.first
	}

	node := &Node{value: v}
	node.next = l.first

	l.first.prev = node

	l.first = node

	return l.first
}

// String is a Stringer interface implementation,
// provides default string representation of list
func (l *List) String() string {
	if l.len == 0 {
		return ""
	}

	parts := []string{}

	head := l.first
	for head != nil {
		parts = append(parts, fmt.Sprintf("%v", head.value))
		head = head.next
	}

	return strings.Join(parts, " ")
}

// Front returns pointer to list head, first item
func (l *List) Front() *Node {
	return l.first
}

// Back returns pointer to list last element
func (l *List) Back() *Node {
	return l.last
}

// Len returns list items count
func (l *List) Len() int {
	return l.len
}
