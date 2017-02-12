package sort

import "testing"

type testQuick struct {
	input []int
}

var testQuickData = []testQuick{
	{
		input: []int{},
	},
	{
		input: []int{2, 1},
	},
	{
		input: []int{10, 20, 5, 4000, 582, 234},
	},
}

func TestQuick(t *testing.T) {
	for _, td := range testQuickData {
		act := Quick(td.input)
		if !IsSorted(act) {
			t.Errorf("sort failed %+v", act)
		}
	}
}
