package helpers

func IsInBounds(row, col int, data [][]int) bool {
	maxRow := len(data) - 1
	maxCol := len(data[0]) - 1

	return row >= 0 && col >= 0 && row <= maxRow && col <= maxCol
}
