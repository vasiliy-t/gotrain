package single

import (
	"fmt"
	"testing"
)

func TestPushBack(t *testing.T) {
	t.Run("Single element ring", func(t *testing.T) {
		t.Run("First node's next points to itself", func(t *testing.T) {
			ring := NewRing()
			ring.PushBack(1)

			if ring.first.Next() != ring.Front() {
				t.Fatalf("First element must point to itself")
			}
		})
	})

	t.Run("Multi element list", func(t *testing.T) {
		ring := NewRing()
		ring.PushBack(0)
		ring.PushBack(1)
		ring.PushBack(2)

		t.Run("Last node's next points to first node", func(t *testing.T) {
			if ring.last.Next() != ring.Front() {
				t.Fatalf("Last element's next must point to first element")
			}
		})

		t.Run("Following each node's next returns to first node", func(t *testing.T) {
			firstNode := ring.Front()
			secondNode := firstNode.Next()
			thirdNode := secondNode.Next()

			if thirdNode.Next() != firstNode {
				t.Fatalf("Following each node's next must return to first node")
			}
		})
	})
}

func TestPushFront(t *testing.T) {
	t.Run("Single element ring", func(t *testing.T) {
		ring := NewRing()
		ring.PushFront(0)

		t.Run("First node's next points to itself", func(t *testing.T) {
			if ring.first.Next() != ring.Front() {
				t.Fatalf("First node's next must point to itself")
			}
		})
	})

	t.Run("Multi element ring", func(t *testing.T) {
		ring := NewRing()
		ring.PushFront(2)
		ring.PushFront(1)
		ring.PushFront(0)

		t.Run("Last node's next points to first node", func(t *testing.T) {
			if ring.last.Next() != ring.Front() {
				t.Fatalf("Last node's next must point to first element")
			}
		})

		t.Run("Following each node's next returns to first node", func(t *testing.T) {
			firstNode := ring.Front()
			secondNode := firstNode.Next()
			thirdNode := secondNode.Next()

			if thirdNode.Next() != firstNode {
				t.Fatalf("Following each node's next must return to first node")
			}
		})
	})
}

func TestIteration(t *testing.T) {
	type testCase struct {
		arrange       func() *Ring
		maxIterations int
	}

	cases := []*testCase{
		{
			arrange: func() *Ring {
				ring := NewRing()
				ring.PushBack(0)

				return ring
			},
			maxIterations: 1000,
		},
		{
			arrange: func() *Ring {
				ring := NewRing()
				ring.PushBack(0)
				ring.PushBack(1)
				ring.PushFront(2)

				return ring
			},
			maxIterations: 1000,
		},
	}

	for _, c := range cases {
		ring := c.arrange()
		t.Run("Infite iteration is possible on "+fmt.Sprintf("%d", ring.Len())+" node ring", func(t *testing.T) {

			ring := NewRing()
			ring.PushBack(0)

			totalIterations := 0
			for e := ring.Front(); e != nil; e = e.Next() {
				if totalIterations >= c.maxIterations {
					break
				}

				totalIterations++
			}

			if totalIterations != c.maxIterations {
				t.Fatalf("Infinite iteration must be possible, but only %d iteration occured", totalIterations)
			}
		})
	}

}
