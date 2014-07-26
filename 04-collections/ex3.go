package strings

func Make2D(source []int, columns int) [][]int {
	matrix := make([][]int, neededRows(source, columns))
	for i, _ := range matrix {
		matrix[i] = make([]int, columns)
	}
	for i, x := range source {
		row := i / columns
		column := i % columns
		matrix[row][column] = x
	}
	return matrix
}

func neededRows(source []int, columns int) int {
	rows := len(source) / columns
	if len(source)%columns != 0 {
		rows++
	}
	return rows
}
