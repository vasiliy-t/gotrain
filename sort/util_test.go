package sort

import "testing"

type isSortedTest struct {
	values []int
	expected bool
}

var isSortedTestData = []isSortedTest{
	{
		values: []int{1,2,3},
		expected: true,
	},
	{
		values: []int{3,1,2},
		expected: false,
	},
}

func TestIsSorted(t *testing.T) {
	for _, td := range isSortedTestData {
		act := IsSorted(td.values)
		if td.expected != act {
			t.Errorf("Failed to assert that %+v is sorted", td.values)
		}
	}
}
