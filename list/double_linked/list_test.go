package double_linked

import (
	"fmt"
	"strings"
	"testing"
)

func TestListLen(t *testing.T) {
	t.Run("Empty list has 0 len", func(t *testing.T) {
		list := NewList()
		actual := list.Len()
		if actual != 0 {
			t.Fatalf("Empty list has non 0 length %d", actual)
		}
	})

	t.Run("PushBack increases len by 1 at a time", func(t *testing.T) {
		list := NewList()

		for i := 1; i < 100; i++ {
			list.PushBack(i)
			actual := list.Len()
			if actual != i {
				t.Fatalf("Unexpected list len after PushBack, expected %d, actual %d", i, actual)
			}
		}
	})

	t.Run("PushFront increases len by 1 at a time", func(t *testing.T) {
		list := NewList()

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
		list := NewList()
		list.PushBack(1)

		onlyNode := list.Front()

		if list.first != onlyNode {
			t.Fatalf("Unexpected first node")
		}
		if list.last != onlyNode {
			t.Fatalf("Unxpected last node")
		}
	})

	t.Run("PushBack creates next links", func(t *testing.T) {
		list := NewList()
		list.PushBack(1)
		list.PushBack(2)

		firstNode := list.Front()
		secondNode := firstNode.Next()

		if firstNode.next != secondNode {
			t.Fatalf("First node must hold next ref to second node")
		}
	})

	t.Run("PushBack creates prev links", func(t *testing.T) {
		list := NewList()
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
		list := NewList()
		list.PushFront(1)

		onlyNode := list.Front()
		if list.first != onlyNode {
			t.Fatalf("Unexpected first node")
		}
		if list.last != onlyNode {
			t.Fatalf("Unxpected last node")
		}
	})

	t.Run("PushFront creates next links", func(t *testing.T) {
		list := NewList()
		list.PushFront(1)
		list.PushFront(2)

		firstNode := list.Front()
		secondNode := firstNode.Next()

		if firstNode.next != secondNode {
			t.Fatalf("First node must hold next ref to second node")
		}
	})

	t.Run("PushFront creates prev links", func(t *testing.T) {
		list := NewList()
		list.PushFront(1)
		list.PushFront(2)

		firstNode := list.Front()
		secondNode := firstNode.Next()

		if secondNode.prev != firstNode {
			t.Fatalf("Second node must hold prev ref to first node")
		}
	})
}

func TestReverse(t *testing.T) {
	t.Run("Reverse rewrites prev and next refs", func(t *testing.T) {
		list := NewList()
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

	t.Run("First node prev is nil after Reverse", func(t *testing.T) {
		list := NewList()
		list.PushBack(0)
		list.PushBack(1)
		list.PushFront(2)

		list.Reverse()
		if list.first.prev != nil {
			t.Fatalf("First node's prev must be nil")
		}
	})

	t.Run("Last node next is nil after Reverse", func(t *testing.T) {
		list := NewList()
		list.PushBack(0)
		list.PushFront(1)
		list.PushBack(3)

		list.Reverse()

		if list.last.next != nil {
			t.Fatalf("Last node's next must be nil")
		}
	})
}

func TestString(t *testing.T) {
	type testCase struct {
		arrange  func() *List
		expected string
	}

	cases := []*testCase{
		{
			arrange: func() *List {
				return NewList()
			},
			expected: "",
		},
		{
			arrange: func() *List {
				ll := NewList()

				ll.PushBack(0)
				ll.PushBack(1)
				ll.PushBack(2)

				return ll
			},
			expected: "0 1 2",
		},
		{
			arrange: func() *List {
				ll := NewList()

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
		list := NewList()
		list.PushFront(0)
		list.PushBack(1)
		list.PushBack(2)

		parts := []string{}
		for e := list.Front(); e != nil; e = e.Next() {
			parts = append(parts, fmt.Sprintf("%v", e.value))
		}

		expected := "0 1 2"
		actual := strings.Join(parts, " ")
		if expected != actual {
			t.Fatalf("Forward iteration must result in elements order first to last, expected %s, actual %s", expected, actual)
		}
	})
	t.Run("Backward iteration results in elements ordering from last to first", func(t *testing.T) {
		list := NewList()
		list.PushBack(0)
		list.PushBack(1)
		list.PushBack(2)

		parts := []string{}

		for e := list.Back(); e != nil; e = e.Prev() {
			parts = append(parts, fmt.Sprintf("%v", e.value))
		}

		expected := "2 1 0"
		actual := strings.Join(parts, " ")
		if expected != actual {
			t.Fatalf("Backward iteration must result in elements order last to first, expected %s, actual %s", expected, actual)
		}
	})
}

func TestSanity(t *testing.T) {
	t.Run("1 item in list", func(t *testing.T) {
		t.Run("next is nil", func(t *testing.T) {
			list := NewList()
			list.PushBack(0)
			onlyNode := list.Front()
			if onlyNode.next != nil {
				t.Fatalf("First node's next must be nil")
			}
		})
		t.Run("prev is nil", func(t *testing.T) {
			list := NewList()
			list.PushBack(0)
			onlyNode := list.Front()
			if onlyNode.next != nil {
				t.Fatalf("First node's prev must be nil")
			}
		})
	})

	t.Run("First and last node doesn't directly linked if more than 2 elements in list", func(t *testing.T) {
		list := NewList()
		list.PushBack(1)
		list.PushFront(0)
		list.PushBack(2)

		if list.first.next == list.last {
			t.Fatalf("first node's next points to last node")
		}

		if list.first.prev == list.last {
			t.Fatalf("first node's prev points to last node")
		}

		if list.last.next == list.first {
			t.Fatalf("last node's next points to first node")
		}

		if list.last.prev == list.first {
			t.Fatalf("last node's prev points to first node")
		}
	})
}
