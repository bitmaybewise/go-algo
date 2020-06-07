package algo

import (
	"testing"

	"github.com/hlmerscher/go-algo/helper"
)

func TestQuickSort(t *testing.T) {
	t.Run("Test Simple Sorting", helper.TestSimpleSorting(QuickSort))
	t.Run("Test Random Collection Sorting", helper.TestSortingRandomCollection(QuickSort))
}

func BenchmarkQuickSortWorstCase(b *testing.B) {
	col := helper.GenInvertedCollectionWithDefault()
	QuickSort(col)
}
