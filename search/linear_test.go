package search

import "testing"

type linSearchData struct {
	haystack []string
	needle   string
	expected int
}

var linSearchTestData = []linSearchData{
	{
		haystack: []string{
			"abc",
			"bcd",
			"def",
		},
		needle:   "qwe",
		expected: -1,
	},
	{
		expected: -1,
	},
	{
		haystack: []string{
			"qwe",
			"Qwe",
		},
		needle:   "Qwe",
		expected: 1,
	},
}

func TestLinearSearch(t *testing.T) {
	for _, item := range linSearchTestData {
		res := Search(item.haystack, item.needle)

		if res != item.expected {
			t.Errorf("failed to assert, expected %d, got %d", item.expected, res)
		}
	}
}
