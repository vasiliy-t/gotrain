package sort

import "math/rand"

func Quick(input []int) []int {
	if len(input) <= 1 {
		return input
	}

	idx := rand.Int31n(int32(len(input)))

	input[0], input[idx] = input[idx], input[0]
	last := 0

	for i := 1; i < len(input); i++ {
		if input[i] < input[0] {
			last++
			input[i], input[last] = input[last], input[i]
		}
	}
	input[0], input[last] = input[last], input[0]

	Quick(input[0:last])
	Quick(input[last:])

	return input
}
