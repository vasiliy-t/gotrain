package eval

import (
	"testing"
)

func TestEval(t *testing.T) {
	type testCase struct {
		expr     string
		expected int
	}

	cases := []testCase{
		{
			expr:     "(2 + 2)",
			expected: 4,
		},
		{
			expr:     "(2 + (2 * 2))",
			expected: 6,
		},
		{
			expr:     "( 0 / (10 - 5))",
			expected: 0,
		},
		{
			expr:     "(s ( 2 + (6 - 4)))",
			expected: 2,
		},
	}

	for _, c := range cases {
		actual := Eval(c.expr)
		if c.expected != actual {
			t.Fatalf("Wrong eval result, expected %d, actual %d", c.expected, actual)
		}
	}
}
