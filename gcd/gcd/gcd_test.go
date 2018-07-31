package gcd

import (
	"testing"
)

type Case struct {
	a, b, expected int
}

var cases = []Case{
	{a: 16, b: 8, expected: 8},
	{a: 72, b: 90, expected: 18},
	{a: 3918848, b: 1653264, expected: 61232},
	{a: 7, b: 3, expected: 1},
}

func TestEqulidGCD(t *testing.T) {
	for _, c := range cases {
		actual := EqulidGCD(c.a, c.b)
		if actual != c.expected {
			t.Fatalf("Expected %d is not equal to actual %d", c.expected, actual)
		}
	}
}
