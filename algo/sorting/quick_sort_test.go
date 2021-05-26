package sorting

import (
	"testing"

	"github.com/hlmerscher/go-algo/misc"
)

func TestQuickSort(t *testing.T) {
	t.Run("Test Simple Sorting", misc.TestSimpleSorting(QuickSort))
	t.Run("Test Random Collection Sorting", misc.TestSortingRandomCollection(QuickSort))
}

func BenchmarkQuickSortWorstCase(b *testing.B) {
	col := misc.GenInvertedCollectionWithDefault()
	QuickSort(col)
}
