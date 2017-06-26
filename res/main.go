package main

import (
	"fmt"
)

func main() {
	c := NewCalc()
	c.Add(1)
	c.Add(2)
	c.Add(42)
	c.Sub(1)
	c.Sub(5)
	fmt.Println(c.Add(10))
}

type Calc struct {
	in  chan func(acc int) int
	out chan int
	acc int
}

func NewCalc() *Calc {
	c := &Calc{
		in:  make(chan func(acc int) int),
		out: make(chan int),
	}

	go c.loop()

	return c
}

func (c *Calc) loop() {
	for v := range c.in {
		c.acc = v(c.acc)
		c.out <- c.acc
	}
}

func (c *Calc) Add(a int) int {
	c.in <- func(acc int) int {
		return acc + a
	}
	return <-c.out
}

func (c *Calc) Sub(a int) int {
	c.in <- func(acc int) int {
		return acc - a
	}
	return <-c.out
}
