package strings

func Flatten(source [][]int) (results []int) {
	for _, slice := range source {
		for _, i := range slice {
			results = append(results, i)
		}
	}
	return results
}
