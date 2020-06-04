package algo

import (
	"math/rand"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	col := []int{7, 5, 1, 8, 3}
	expected := []int{1, 3, 5, 7, 8}
	SelectionSort(col)

	for i := 0; i < len(col); i++ {
		if col[i] != expected[i] {
			t.Errorf("Value %d at pos %d should be %d", col[i], i, expected[i])
		}
	}
}

func TestSelectionSortRandomCollection(t *testing.T) {
	col := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		col = append(col, rand.Intn(10000))
	}

	SelectionSort(col)

	for i := 0; i < 999; i++ {
		nextIdx := i + 1
		if col[i] > col[nextIdx] {
			t.Errorf("Value %d at index %d is bigger than value %d at index %d", col[i], i, col[nextIdx], nextIdx)
		}
	}
}
