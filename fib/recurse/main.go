package main

import (
	"fmt"
)

func calculate(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	return calculate(n-1) + calculate(n-2)
}

func main() {
	fmt.Println(calculate(6))
}
