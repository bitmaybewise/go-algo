package helper

import "math/rand"

// GenRandomCollection generates random collection of integers up to N size
func GenRandomCollection(n int) []int {
	col := make([]int, n)
	for i := 0; i < n; i++ {
		col[i] = rand.Intn(n)
	}
	return col
}

// IsSorted checks if collection of integers is sorted
func IsSorted(col []int) bool {
	size := len(col)
	for i := 0; i < size-1; i++ {
		nextIdx := i + 1
		if col[i] > col[nextIdx] {
			return false
		}
	}
	return true
}
