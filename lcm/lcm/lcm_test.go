package lcm

import (
	"testing"
)

type Case struct {
	a, b, expected int
}

var cases = []Case{
	{a: 3, b: 7, expected: 21},
}

func TestLCM(t *testing.T) {
	for _, c := range cases {
		actual := LCM(c.a, c.b)

		if c.expected != actual {
			t.Fatalf("Expected %d not equal to actual %d", c.expected, actual)
		}
	}
}
