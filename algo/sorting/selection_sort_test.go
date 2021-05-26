package sorting

import (
	"testing"

	"github.com/hlmerscher/go-algo/misc"
)

func TestSelectionSort(t *testing.T) {
	t.Run("Test Simple Sorting", misc.TestSimpleSorting(SelectionSort))
	t.Run("Test Random Collection Sorting", misc.TestSortingRandomCollection(SelectionSort))
}

func BenchmarkSelectionSortWorstCase(b *testing.B) {
	col := misc.GenInvertedCollectionWithDefault()
	SelectionSort(col)
}
