package linkedlist

import "testing"

func TestIsEmpty(t *testing.T) {
	t.Run("Empty stack", func(t *testing.T) {
		s := NewStack()
		expected := true
		actual := s.IsEmpty()
		if expected != actual {
			t.Fatalf("IsEmpty must return true for empty stack, got %t", actual)
		}
	})

	t.Run("Empty stack, push one element", func(t *testing.T) {
		s := NewStack()
		s.Push(1)

		expected := false
		actual := s.IsEmpty()

		if expected != actual {
			t.Fatalf("IsEmpty must return false for stack containing 1 element, got %t", actual)
		}
	})
	t.Run("Pop all elements from stack", func(t *testing.T) {
		s := NewStack()
		s.Push(1)
		s.Pop()

		expected := true
		actual := s.IsEmpty()

		if expected != actual {
			t.Fatalf("IsEmpty must return true after all elements popped from stack, got %t", actual)
		}
	})
}

func TestSanity(t *testing.T) {
	t.Run("Pop returns most recently pushed element", func(t *testing.T) {
		s := NewStack()
		expected := "qwerty"
		s.Push(0)
		s.Push(expected)

		actual := s.Pop()
		if expected != actual {
			t.Fatalf("Pop returned not recently pushed element, expected %s, got %s", expected, actual)
		}
	})

	t.Run("Pop returns elements in order, from most recent to least recent", func(t *testing.T) {
		s := NewStack()

		first := 1
		second := "qwerty"
		third := true

		s.Push(first)
		s.Push(second)
		s.Push(third)

		actualThird := s.Pop()
		actualSecond := s.Pop()
		actualFirst := s.Pop()

		if first != actualFirst || second != actualSecond || third != actualThird {
			t.Fatalf("Pop order messed, expected [%d, %s, %t], actual [%d, %s, %t]", first, second, third, actualFirst, actualSecond, actualThird)
		}
	})
}
