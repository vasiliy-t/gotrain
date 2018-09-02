package linkedlist

// ListBackedStack is an implementation of stack backed by simple linked list
type ListBackedStack struct {
	first *Node
}

// NewStack initializes and returns new ListBackedStack instance
func NewStack() *ListBackedStack {
	return &ListBackedStack{}
}

// Node is a list item implementation
type Node struct {
	value interface{}
	next  *Node
}

// IsEmpty return boolean identifying is stack empty or not
func (s *ListBackedStack) IsEmpty() bool {
	return s.first == nil
}

// Pop returns most recently pushed element
func (s *ListBackedStack) Pop() interface{} {
	val := s.first.value
	s.first = s.first.next
	return val
}

// Push add element to the front of stack
func (s *ListBackedStack) Push(item interface{}) *Node {
	n := &Node{value: item}
	n.next = s.first
	s.first = n
	return n
}
