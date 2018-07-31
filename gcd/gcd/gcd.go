package gcd

// EqulidGCD is an implementation of Equlid greatest common divisor algorithm
func EqulidGCD(a, b int) int {
	if b == 0 {
		return a
	}

	aPrime := a % b

	return EqulidGCD(b, aPrime)
}
