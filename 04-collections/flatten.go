package strings

// Takes a slice of slices of integers, returns a single slice of all the integers concatenated together
func Flatten(source [][]int) (results []int) {
	for _, slice := range source {
		for _, i := range slice {
			results = append(results, i)
		}
	}
	return
}
