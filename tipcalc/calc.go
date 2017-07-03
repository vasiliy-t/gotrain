package main

func Calculate(amount float64, rate float64) float64 {
	tip := amount * rate / 100
	return tip
}
