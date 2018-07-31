package main

import (
	"flag"
	"fmt"

	"github.com/vasiliy-t/gotrain/gcd/gcd"
)

func main() {
	a := flag.Int("a", 0, "a")
	b := flag.Int("b", 0, "b")

	flag.Parse()

	gcd := gcd.EqulidGCD(*a, *b)
	fmt.Printf("GCD(a, b) = %d\n", gcd)
}
