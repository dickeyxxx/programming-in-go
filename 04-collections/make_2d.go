package strings

// Given a slice of integers, returns a matrix of integers with columns in each.
func Make2D(source []int, columns int) (matrix [][]int) {
	matrix = make([][]int, neededRows(source, columns))
	for i, _ := range matrix {
		matrix[i] = make([]int, columns)
	}
	for i, x := range source {
		row := i / columns
		column := i % columns
		matrix[row][column] = x
	}
	return
}

func neededRows(source []int, columns int) (rows int) {
	rows = len(source) / columns
	if len(source)%columns != 0 {
		rows++
	}
	return
}
