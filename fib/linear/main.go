package main

import (
	"fmt"
)

func calculate(n int) int {
	alloc := []int{}
	alloc = append(alloc, []int{0, 1}...)

	for i := 2; i <= n; i++ {
		alloc = append(alloc, alloc[i-1]+alloc[i-2])
	}

	return alloc[n]
}

func main() {
	fmt.Println(calculate(1))
}
