package search

func Search(haystack []string, needle string) int {
	for k, v := range haystack {
		if v == needle {
			return k
		}
	}
	return -1
}
