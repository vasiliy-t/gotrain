package linked

import (
	"fmt"
	"strings"
)

// LinkedList is a container structer for list items
type LinkedList struct {
	head *Node
	tail *Node
	len int
}

// Node is an implementation of list item
type Node struct {
	value interface{}
	next *Node
}

// Next returns next item or nil if no exists
func (n *Node) Next() *Node {
	return n.next
}

// NewLinkedList returns new LinkedList instance
func NewLinkedList() *LinkedList {
	ll := &LinkedList{}
	return ll
}

// Reverse reverses linked list items
// Reverse is mutating operation
func (l *LinkedList) Reverse() *LinkedList {
	n := l.head

	var prev *Node
	var next *Node

	for n != nil {
		next = n.next
		n.next = prev

		prev = n
		n = next
	}

	h := l.head
	t := l.tail
	

	l.head = t
	l.tail = h

	return l
}

// PushBack pushes new item to list end
func (l *LinkedList) PushBack(v interface{}) *Node {
	l.len = l.len + 1

	if l.head == nil {
		l.head = &Node{value: v}
		l.tail = l.head

		return l.head
	}

	node := &Node{value: v}
	l.tail.next = node
	l.tail = node 

	return node
}

// PushFront add new item in front of all items
func (l *LinkedList) PushFront(v interface{}) *Node {
	l.len = l.len + 1

	if l.head == nil {
		l.head = &Node{value: v}
		l.tail = l.head

		return l.head
	}

	node := &Node{value: v}
	node.next = l.head
	l.head = node

	return l.head
}

// String is a Stringer interface implementation,
// provides default string representation of list
func (l *LinkedList) String() string {
	if l.len == 0 {
		return ""
	}

	parts := []string{}
	
	head := l.head 
	for head != nil {
		parts = append(parts, fmt.Sprintf("%v",head.value))
		head = head.next	
	}
	
	return strings.Join(parts, " ")
}

// Front returns pointer to list head, first item
func (l *LinkedList) Front() *Node {
	return l.head
}

// Len returns list items count
func (l *LinkedList) Len() int {
	return l.len
}
