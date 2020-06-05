package algo

import (
	"testing"

	"github.com/hlmerscher/go-algo/helper"
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

func TestSelectionSortWithBigRandomCollection(t *testing.T) {
	col := helper.GenRandomCollection(10000)

	if SelectionSort(col); !helper.IsSorted(col) {
		t.Error("Collection is not sorted")
	}
}

func BenchmarkSelectionSortWorstCase(b *testing.B) {
	col := helper.GenInvertedCollectionWithDefault()
	SelectionSort(col)
}
