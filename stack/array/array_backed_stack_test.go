package array

import (
	"fmt"
	"testing"
)

func TestPush(t *testing.T) {
	t.Run("Push increments Idx", func(t *testing.T) {
		s := NewStack()
		s.Push("str1")
		s.Push("str2")
		s.Push("str3")

		if s.(*FixedSizeStack).Idx != 3 {
			t.Fatalf("Idx must increment on Push, expected 2, actual %d", s.(*FixedSizeStack).Idx)
		}
	})
}

func TestPop(t *testing.T) {
	t.Run("Pop decrements Idx", func(t *testing.T) {
		s := NewStack()
		s.Push("str1")
		s.Push("str2")

		expected := "str2"
		actual := s.Pop()
		if expected != actual {
			t.Fatalf("Pop return value doesn't match, expected %s, actual %s", expected, actual)
		}
	})
}

func TestSanity(t *testing.T) {
	t.Run("Pop, Push series results in expected stack contents", func(t *testing.T) {
		s := NewStack()
		s.Push("str1")
		s.Push("str2")
		s.Pop()
		s.Push("str3")

		if s.(*FixedSizeStack).Idx != 2 {
			t.Fatalf("Idx must equal 3")
		}

		if s.Pop() != "str3" {
			t.Fatalf("Most recent element must be str3")
		}
	})

	t.Run("Stack overflow panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Fatalf("Stack overflow must panic")
			}
		}()
		s := NewStack()
		for i := 0; i < 10; i++ {
			s.Push(fmt.Sprintf("%d", i))
		}

		s.Push("qwerty")
	})
}
