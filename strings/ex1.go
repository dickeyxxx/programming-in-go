package strings

func UniqueInts(source []int) (results []int) {
	seen := map[int]bool{}
	for _, i := range source {
		if _, found := seen[i]; !found {
			results = append(results, i)
			seen[i] = true
		}
	}
	return results
}
