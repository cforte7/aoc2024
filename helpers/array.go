package helpers

func IsInBounds[T any](row, col int, data [][]T) bool {
	maxRow := len(data) - 1
	maxCol := len(data[0]) - 1

	return row >= 0 && col >= 0 && row <= maxRow && col <= maxCol
}

func ReplaceInArray[T any](incoming []T, new []T, index int) []T {
	out := make([]T, 0)
	out = append(out, incoming[:index]...)
	out = append(out, new...)
	out = append(out, incoming[index+1:]...)
	return incoming
}

func MapToSlice[K comparable, V any](incoming map[K]V) []K {
	out := make([]K, 0)

	for k := range incoming {
		out = append(out, k)
	}
	return out
}
