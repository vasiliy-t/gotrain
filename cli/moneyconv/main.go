package main

import (
	"flag"
	"fmt"
)

func main() {
	var amount, rate, res int

	flag.IntVar(&amount, "amount", 0, "Amount of currency to convert")
	flag.IntVar(&rate, "rate", 0, "Exchange rate")
	flag.Parse()

	res = amount * rate
	fmt.Printf("%d\n", res)
}
