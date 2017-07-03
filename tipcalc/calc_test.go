package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculate_NonZeroRate_ReturnsExpectedTip(t *testing.T) {
	testData := []struct {
		amount   float64
		rate     float64
		expected float64
	}{
		{
			0,
			0,
			0,
		},
		{
			1,
			10,
			0.1,
		},
	}

	for _, v := range testData {
		assert.Equal(t, v.expected, Calculate(v.amount, v.rate))
	}
}
