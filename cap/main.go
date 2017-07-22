package main

import (
	"flag"
	"fmt"
	"math"
)

func main() {
	s := flag.Float64("interest", 0.00, "Interest rate")
	x := flag.Float64("amount", 0.00, "Cash amount")
	m := flag.Float64("months", 0, "Deposit time")
	flag.Parse()

	res := *x * math.Pow((1 + *s/100), *m/12)
	fmt.Printf("Resulting amount: %f", res)
}
