package sort

func IsSorted(input []int) bool {
	n := len(input)
	for i := 0; i < n -1; i++ {
		if input[i] > input[i + 1] {
			return false
		}
	}

	return true
}
