package queue

// Queue is an queue implementation
type Queue struct {
	first, last *Node

	size int
}

// NewQueue initializes and return new Queue instances
func NewQueue() *Queue {
	return &Queue{}
}

// IsEmpty returns boolean indicating if queue is empty
// Returns true if there is no items in queue
func (q *Queue) IsEmpty() bool {
	return q.first == nil
}

// Enqueue push element to queue
func (q *Queue) Enqueue(item interface{}) *Node {
	defer func() { q.size++ }()
	n := &Node{value: item}

	if q.IsEmpty() {
		q.first = n
		q.last = n

		return n
	}

	q.last.next = n
	q.last = n

	return n
}

// Dequeue return least recent added item value
func (q *Queue) Dequeue() interface{} {
	defer func() {
		if !q.IsEmpty() || q.size != 0 {
			q.size--
		}
	}()

	if q.IsEmpty() {
		return nil
	}

	val := q.first.value

	q.first = q.first.next

	if q.IsEmpty() {
		q.last = nil
	}

	return val
}

// Size return number of items in queue
func (q *Queue) Size() int {
	return q.size
}

// Node is an queue item container
type Node struct {
	value interface{}
	next  *Node
}
