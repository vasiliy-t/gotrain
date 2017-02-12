package search

func BinarySearch(haystack []int, needle int) int {
	var low, high, mid int

	high = len(haystack) - 1

	for low <= high {
		mid = (low + high) / 2

		e := haystack[mid]

		if needle < e {
			high = mid - 1
		}

		if needle > e {
			low = mid + 1
		}

		if needle == e {
			return mid
		}
	}

	return -1
}
