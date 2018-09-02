package queue

import "testing"

func TestIsEmpty(t *testing.T) {
	t.Run("Empty queue", func(t *testing.T) {
		q := NewQueue()

		expected := true
		actual := q.IsEmpty()

		if expected != actual {
			t.Fatalf("IsEmpty must return %t for empty queue, actual %t", expected, actual)
		}
	})
	t.Run("Non empty queue", func(t *testing.T) {
		q := NewQueue()
		q.Enqueue(1)

		expected := false
		actual := q.IsEmpty()

		if expected != actual {
			t.Fatalf("IsEmpty must return %t for non empty queue, actual %t", expected, actual)
		}
	})
}

func TestEnqueue(t *testing.T) {
	t.Run("Enqueue to empty queue", func(t *testing.T) {
		q := NewQueue()
		expectedValue := 1
		q.Enqueue(expectedValue)

		if q.last.value != expectedValue || q.first.value != expectedValue {
			t.Fatalf("First element %d must be equal to last element %d must be equal to expected %d", q.first.value, q.last.value, expectedValue)
		}
	})

	t.Run("Enqueue multiple elements", func(t *testing.T) {
		q := NewQueue()

		expectedFirst := 1
		expectedLast := 2

		q.Enqueue(expectedFirst)
		q.Enqueue(expectedLast)

		if q.first.value != expectedFirst {
			t.Fatalf("Queue first value must be %d, got %d", expectedFirst, q.first.value)
		}
		if q.last.value != expectedLast {
			t.Fatalf("Queue last value must be %d, got %d", expectedLast, q.last.value)
		}
	})

	t.Run("Each queue item points to next item", func(t *testing.T) {
		q := NewQueue()

		first := q.Enqueue(1)
		second := q.Enqueue(2)
		third := q.Enqueue(3)

		if first.next != second {
			t.Fatalf("First queue item doesn't point to second item expected %p, got %p", second, first.next)
		}

		if second.next != third {
			t.Fatalf("Second queue item must point to third item, expected %p, got %p", third, second.next)
		}
	})
}

func TestDequeue(t *testing.T) {
	t.Run("Empty queue", func(t *testing.T) {
		q := NewQueue()
		actual := q.Dequeue()

		if actual != nil {
			t.Fatalf("Dequeue must return nil, got %v", actual)
		}
	})

	t.Run("Single item in queue", func(t *testing.T) {
		q := NewQueue()
		onlyNode := q.Enqueue(1)
		val := q.Dequeue()

		if onlyNode.value != val {
			t.Fatalf("Must return only item value, expected %d, got %d", onlyNode.value, val)
		}

		t.Run("Dequeued item removed from queue", func(t *testing.T) {
			q := NewQueue()
			q.Enqueue(1)
			q.Dequeue()

			if q.first != nil || q.last != nil {
				t.Fatalf("Queue first and last elements must be nil, got %p %p", q.first, q.last)
			}
		})
	})

	t.Run("Multiple elements in queue", func(t *testing.T) {
		t.Run("Dequeue returns least recent enqueued item", func(t *testing.T) {
			q := NewQueue()
			first := q.Enqueue(1)
			second := q.Enqueue(2)
			third := q.Enqueue(3)

			val := q.Dequeue()
			if val != first.value {
				t.Fatalf("First: Must return least recent item, expected %d, actual %d", first.value, val)
			}

			val = q.Dequeue()
			if val != second.value {
				t.Fatalf("Second: Must return least recent item, expected %d, actual %d", first.value, val)
			}

			val = q.Dequeue()
			if val != third.value {
				t.Fatalf("Third: Must return least recent item, expected %d, actual %d", third.value, val)
			}
		})

		t.Run("All elements removed", func(t *testing.T) {
			q := NewQueue()
			q.Enqueue(1)
			q.Enqueue("qwerty")
			q.Enqueue(2)

			q.Dequeue()
			q.Dequeue()
			q.Dequeue()

			if q.first != nil || q.last != nil {
				t.Fatalf("Queue must be empty, got first %p, last %p", q.first, q.last)
			}
		})
	})
}

func TestSize(t *testing.T) {
	t.Run("Empty queue size is 0", func(t *testing.T) {
		q := NewQueue()

		expected := 0
		actual := q.Size()

		if expected != actual {
			t.Fatalf("Empty queue Size must return nil, expected %d, actual %d", expected, actual)
		}
	})

	t.Run("Each Enqueue call increments size by 1", func(t *testing.T) {
		q := NewQueue()

		for i := 1; i < 100; i++ {
			q.Enqueue(i)

			actual := q.Size()
			if actual != i {
				t.Fatalf("Queue size must increment on each Enqueue call, expected %d, actual %d", i, actual)
			}
		}
	})

	t.Run("Each Dequeue call decrements size by 1", func(t *testing.T) {
		q := NewQueue()

		for i := 1; i <= 100; i++ {
			q.Enqueue(i)

			actual := q.Size()
			if actual != i {
				t.Fatalf("Queue size must increment on each Enqueue call, expected %d, actual %d", i, actual)
			}
		}

		for i := 100; i > 0; i-- {
			q.Dequeue()
			actual := q.Size()
			expected := i - 1

			if actual != expected {
				t.Fatalf("Queue size must be decremented on each Dequeue call, expected %d, actual %d", expected, actual)
			}
		}
	})

	t.Run("Size can't be negative", func(t *testing.T) {
		q := NewQueue()
		for i := 100; i > 0; i-- {
			q.Dequeue()
			actual := q.Size()

			if actual != 0 {
				t.Fatalf("Queue size must be 0, expected %d, actual %d", 0, actual)
			}
		}
	})
}
