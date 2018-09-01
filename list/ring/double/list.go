package double

import (
	"fmt"
	"strings"
)

// Ring is a container struct for list items
type Ring struct {
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

// NewRing returns new Ring instance
func NewRing() *Ring {
	ll := &Ring{}
	return ll
}

// Reverse reverses list items
// Reverse is mutating operation
func (l *Ring) Reverse() *Ring {
	n := l.first

	var prev *Node
	var next *Node

	maxIterations := l.Len()
	totalIterations := 0

	for n != nil {
		if totalIterations >= maxIterations {
			break
		}
		next = n.next

		n.next = prev
		n.prev = next

		prev = n
		n = next
		totalIterations++
	}

	h := l.first
	t := l.last

	l.first = t
	l.last = h

	l.first.prev = l.last
	l.last.next = l.first

	return l
}

// PushBack pushes new item to list end
func (l *Ring) PushBack(v interface{}) *Node {
	l.len = l.len + 1

	if l.first == nil {
		l.first = &Node{value: v}
		l.last = l.first

		l.first.next = l.first
		l.first.prev = l.first

		l.last.next = l.last
		l.first.prev = l.last

		return l.first
	}

	node := &Node{value: v}
	l.last.next = node

	node.prev = l.last

	l.last = node

	l.first.prev = l.last
	l.last.next = l.first

	return node
}

// PushFront add new item in front of all items
func (l *Ring) PushFront(v interface{}) *Node {
	l.len = l.len + 1

	if l.first == nil {
		l.first = &Node{value: v}
		l.last = l.first

		l.first.next = l.first
		l.first.prev = l.first

		l.last.next = l.last
		l.first.prev = l.last

		return l.first
	}

	node := &Node{value: v}
	node.next = l.first

	l.first.prev = node

	l.first = node
	l.first.prev = l.last
	l.last.next = l.first

	return l.first
}

// String is a Stringer interface implementation,
// provides default string representation of list
func (l *Ring) String() string {
	if l.len == 0 {
		return ""
	}

	parts := []string{}

	head := l.first
	maxIterations := l.Len()
	totalIterations := 0
	for head != nil {
		if totalIterations >= maxIterations {
			break
		}
		parts = append(parts, fmt.Sprintf("%v", head.value))
		head = head.next
		totalIterations++
	}

	return strings.Join(parts, " ")
}

// Front returns pointer to list head, first item
func (l *Ring) Front() *Node {
	return l.first
}

// Back returns pointer to list last element
func (l *Ring) Back() *Node {
	return l.last
}

// Len returns list items count
func (l *Ring) Len() int {
	return l.len
}
