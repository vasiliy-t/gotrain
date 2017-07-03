package main

import (
	"flag"
	"fmt"
)

func main() {
	amount := flag.Float64("amount", 10.00, "Amount")
	rate := flag.Float64("rate", 10, "Rate")

	flag.Parse()

	fmt.Println(Calculate(*amount, *rate))
}
