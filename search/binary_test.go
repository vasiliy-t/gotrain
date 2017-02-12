package search

import (
	"testing"
)

type binSearchData struct {
	haystack []int
	needle   int
	expected int
}

var binSearchTestData = []binSearchData{
	{
		haystack: []int{
			0, 1, 2, 3,
		},
		expected: -1,
		needle:   10,
	},
	{
		haystack: []int{
			0, 1, 2, 3, 4, 5,
		},
		expected: 5,
		needle:   5,
	},
	{
		haystack: []int{
			10, 20, 30, 40, 50, 60, 100, 5900,
		},
		needle:   100,
		expected: 6,
	},
	{
		haystack: []int{},
		needle:   10,
		expected: -1,
	},
}

func TestBinarySearch(t *testing.T) {
	for _, v := range binSearchTestData {
		res := BinarySearch(v.haystack, v.needle)

		if res != v.expected {
			t.Errorf("failed to assert that actual %d equals expected %d", res, v.expected)
		}
	}
}
