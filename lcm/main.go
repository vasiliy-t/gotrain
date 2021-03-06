package main

import (
	"flag"
	"fmt"

	"github.com/vasiliy-t/gotrain/lcm/lcm"
)

func main() {
	a := flag.Int("a", 0, "a")
	b := flag.Int("b", 0, "b")

	flag.Parse()

	fmt.Printf("LCM(a, b) = %d\n", lcm.LCM(*a, *b))
}
