package lcm

import "github.com/vasiliy-t/gotrain/gcd/gcd"

// LCM is an implementation of least common multiplier algorithm
func LCM(a, b int) int {
	return a * b / gcd.EqulidGCD(a, b)
}
