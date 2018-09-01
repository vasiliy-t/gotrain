package linked

import (
	"fmt"
	"strings"
	"testing"
)

func TestIteration_ListHasItems_IterateOneByOneInOrder(t *testing.T) {
	ll := NewLinkedList()

	ll.PushBack(0)
	ll.PushBack(1)
	ll.PushBack(2)

	expected := "0 1 2"
	parts := []string{}

	for e := ll.Front(); e != nil; e = e.Next() {
		parts = append(parts, fmt.Sprintf("%v", e.value))
	}

	actual := strings.Join(parts, " ")

	if expected != actual {
		t.Fatalf("Failed to iterate, expected %s, actual %s", expected, actual)
	}
}

func TestReverse_ReversesListItems(t *testing.T) {
	ll := NewLinkedList()

	ll.PushBack(1)
	ll.PushBack(2)
	ll.PushBack(3)

	ll.Reverse()

	expected := "3 2 1"
	actual := fmt.Sprintf("%s", ll)

	if expected != actual {
		t.Fatalf("failed to reverse list, expected %s, actual %s", expected, actual)
	}
}

func TestIntegration_Reverse_Push(t *testing.T) {
	ll := NewLinkedList()

	ll.PushFront(0)
	ll.PushBack(1)
	ll.Reverse()

	ll.PushBack(2)
	ll.PushFront(3)

	expected := "3 1 0 2"
	actual := fmt.Sprintf("%s", ll)

	if expected != actual {
		t.Fatalf("integration failed, expected %s, actual %s", expected, actual)
	}
}

func TestLen_ReturnsActualListItemsCount(t *testing.T) {
	type testCase struct {
		arrange  func() *List
		expected int
	}

	testCases := []testCase{
		{
			arrange: func() *List {
				ll := NewLinkedList()
				ll.PushFront(1)
				ll.PushBack(2)
				ll.PushFront(3)
				return ll
			},
			expected: 3,
		},
		{
			arrange: func() *List {
				ll := NewLinkedList()
				return ll
			},
			expected: 0,
		},
	}

	for _, c := range testCases {
		ll := c.arrange()
		actual := ll.Len()
		if actual != c.expected {
			t.Fatalf("Item count mismatch, expected %d, actual %d", c.expected, actual)
		}
	}
}

func TestToString_LinkedtListWithData_ReturnsStringOfValuesInOrder(t *testing.T) {
	list := NewLinkedList()

	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)

	expected := "1 2 3"
	actual := fmt.Sprintf("%s", list)

	if actual != expected {
		t.Fatalf("Got unexpected result, actual %s, expected %s", actual, expected)
	}
}

func TestPushFront_AddElementsToLinkedListFront(t *testing.T) {
	list := NewLinkedList()

	list.PushFront(3)
	list.PushFront(2)
	list.PushFront(1)

	expected := "1 2 3"
	actual := fmt.Sprintf("%s", list)

	if actual != expected {
		t.Fatalf("Got unexpected result, actual %s, expected %s", actual, expected)
	}
}

func TestToString_EmptyLinkedList_ReturnsEmptyString(t *testing.T) {
	list := NewLinkedList()

	expected := ""
	actual := fmt.Sprintf("%s", list)

	if actual != expected {
		t.Fatalf("Empty list ToString must return empty string, got %v", actual)
	}
}
