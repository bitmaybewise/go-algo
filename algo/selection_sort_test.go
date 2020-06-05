package algo

import (
	"testing"

	"github.com/hlmerscher/go-algo/helper"
)

func TestSelectionSort(t *testing.T) {
	t.Run("Test Simple Sorting", helper.TestSimpleSorting(SelectionSort))
	t.Run("Test Random Collection Sorting", helper.TestSortingRandomCollection(SelectionSort))
}

func BenchmarkSelectionSortWorstCase(b *testing.B) {
	col := helper.GenInvertedCollectionWithDefault()
	SelectionSort(col)
}
