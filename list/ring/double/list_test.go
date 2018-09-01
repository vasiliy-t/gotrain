package double

import (
	"fmt"
	"strings"
	"testing"
)

func TestListLen(t *testing.T) {
	t.Run("Empty list has 0 len", func(t *testing.T) {
		list := NewRing()
		actual := list.Len()
		if actual != 0 {
			t.Fatalf("Empty list has non 0 length %d", actual)
		}
	})

	t.Run("PushBack increases len by 1 at a time", func(t *testing.T) {
		list := NewRing()

		for i := 1; i < 100; i++ {
			list.PushBack(i)
			actual := list.Len()
			if actual != i {
				t.Fatalf("Unexpected list len after PushBack, expected %d, actual %d", i, actual)
			}
		}
	})

	t.Run("PushFront increases len by 1 at a time", func(t *testing.T) {
		list := NewRing()

		for i := 1; i < 100; i++ {
			list.PushFront(i)
			actual := list.Len()
			if actual != i {
				t.Fatalf("Unexpected list len after PushFront, expected %d, actual %d", i, actual)
			}
		}
	})
}

func TestPushBack(t *testing.T) {
	t.Run("PushBack to empty list set node as first and last", func(t *testing.T) {
		ring := NewRing()
		ring.PushBack(1)

		onlyNode := ring.Front()

		if ring.first != onlyNode {
			t.Fatalf("Unexpected first node")
		}
		if ring.last != onlyNode {
			t.Fatalf("Unxpected last node")
		}

		t.Run("First node's next points to itselft", func(t *testing.T) {
			firstNode := ring.Front()
			if firstNode.Next() != firstNode {
				t.Fatalf("First node must point to itself")
			}
		})
	})

	t.Run("PushBack creates next links", func(t *testing.T) {
		list := NewRing()
		list.PushBack(1)
		list.PushBack(2)

		firstNode := list.Front()
		secondNode := firstNode.Next()

		if firstNode.next != secondNode {
			t.Fatalf("First node must hold next ref to second node")
		}
	})

	t.Run("PushBack creates prev links", func(t *testing.T) {
		list := NewRing()
		list.PushBack(0)
		list.PushBack(1)

		firstNode := list.Front()
		secondNode := firstNode.Next()

		if secondNode.prev != firstNode {
			t.Fatalf("Second node must hold prev ref to first node")
		}
	})
}

func TestPushFront(t *testing.T) {
	t.Run("PushFront to empty list sets node as first and last", func(t *testing.T) {
		ring := NewRing()
		ring.PushFront(1)

		onlyNode := ring.Front()
		if ring.first != onlyNode {
			t.Fatalf("Unexpected first node")
		}
		if ring.last != onlyNode {
			t.Fatalf("Unxpected last node")
		}

		t.Run("First node's next points to itselft", func(t *testing.T) {
			firstNode := ring.Front()
			if firstNode.Next() != firstNode {
				t.Fatalf("First node must point to itself")
			}
		})
	})

	t.Run("PushFront creates next links", func(t *testing.T) {
		list := NewRing()
		list.PushFront(1)
		list.PushFront(2)

		firstNode := list.Front()
		secondNode := firstNode.Next()

		if firstNode.next != secondNode {
			t.Fatalf("First node must hold next ref to second node")
		}

		t.Run("Last node's next points to first node", func(t *testing.T) {
			if secondNode.Next() != firstNode {
				t.Fatalf("Last node's next must point to first node")
			}
		})
	})

	t.Run("PushFront creates prev links", func(t *testing.T) {
		list := NewRing()
		list.PushFront(1)
		list.PushFront(2)

		firstNode := list.Front()
		secondNode := firstNode.Next()

		if secondNode.Prev() != firstNode {
			t.Fatalf("Second node must hold prev ref to first node")
		}

		t.Run("First node's prev points to last node", func(t *testing.T) {
			if firstNode.Prev() != secondNode {
				t.Fatalf("First node's prev must point to last node")
			}
		})
	})
}

func TestReverse(t *testing.T) {
	t.Run("Reverse rewrites prev and next refs", func(t *testing.T) {
		list := NewRing()
		list.PushBack(0)
		list.PushBack(1)

		list.Reverse()

		firstNode := list.Front()
		secondNode := firstNode.Next()

		if firstNode.next != secondNode {
			t.Fatalf("First node's next must point to second node")
		}

		if secondNode.prev != firstNode {
			t.Fatalf("Second node's prev must point to first node")
		}
	})

	t.Run("First node prev points to last node after Reverse", func(t *testing.T) {
		ring := NewRing()
		ring.PushBack(0)
		ring.PushBack(1)
		ring.PushFront(2)

		ring.Reverse()

		if ring.first.Prev() != ring.last {
			t.Fatalf("First node's prev must point to last node")
		}
	})

	t.Run("Last node next is not nil after Reverse", func(t *testing.T) {
		ring := NewRing()
		ring.PushBack(0)
		ring.PushFront(1)
		ring.PushBack(3)

		ring.Reverse()

		if ring.last.Next() != ring.Front() {
			t.Fatalf("Last node's next must point to first node")
		}
	})
}

func TestString(t *testing.T) {
	type testCase struct {
		arrange  func() *Ring
		expected string
	}

	cases := []*testCase{
		{
			arrange: func() *Ring {
				return NewRing()
			},
			expected: "",
		},
		{
			arrange: func() *Ring {
				ll := NewRing()

				ll.PushBack(0)
				ll.PushBack(1)
				ll.PushBack(2)

				return ll
			},
			expected: "0 1 2",
		},
		{
			arrange: func() *Ring {
				ll := NewRing()

				ll.PushBack(1)
				ll.PushFront(0)
				ll.PushBack(2)

				ll.Reverse()
				ll.Reverse()
				ll.Reverse()

				return ll
			},
			expected: "2 1 0",
		},
	}

	for _, c := range cases {
		l := c.arrange()
		actual := fmt.Sprintf("%s", l)
		if c.expected != actual {
			t.Fatalf("String representation of List doesn't match, expected %s, actual %s", c.expected, actual)
		}
	}
}

func TestIteration(t *testing.T) {
	t.Run("Forward iteration results in elements ordering from first to last", func(t *testing.T) {
		list := NewRing()
		list.PushFront(0)
		list.PushBack(1)
		list.PushBack(2)

		parts := []string{}
		maxIterations := 3
		totalIterations := 0

		for e := list.Front(); e != nil; e = e.Next() {
			if totalIterations >= maxIterations {
				break
			}
			parts = append(parts, fmt.Sprintf("%v", e.value))
			totalIterations++
		}

		expected := "0 1 2"
		actual := strings.Join(parts, " ")
		if expected != actual {
			t.Fatalf("Forward iteration must result in elements order first to last, expected %s, actual %s", expected, actual)
		}
	})
	t.Run("Backward iteration results in elements ordering from last to first", func(t *testing.T) {
		list := NewRing()
		list.PushBack(0)
		list.PushBack(1)
		list.PushBack(2)

		parts := []string{}
		maxIterations := 3
		totalIterations := 0

		for e := list.Back(); e != nil; e = e.Prev() {
			if totalIterations >= maxIterations {
				break
			}
			parts = append(parts, fmt.Sprintf("%v", e.value))
			totalIterations++
		}

		expected := "2 1 0"
		actual := strings.Join(parts, " ")
		if expected != actual {
			t.Fatalf("Backward iteration must result in elements order last to first, expected %s, actual %s", expected, actual)
		}
	})
}

func TestSanity(t *testing.T) {
	t.Run("First and last node directly linked if more than 2 elements in list", func(t *testing.T) {
		ring := NewRing()
		ring.PushBack(1)
		ring.PushFront(0)
		ring.PushBack(2)

		if ring.Front().Prev() != ring.Back() {
			t.Fatalf("first node's prev must point to last node")
		}

		if ring.Back().Next() != ring.Front() {
			t.Fatalf("last node's next must point to first node")
		}
	})
}
