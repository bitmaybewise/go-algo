package sorting

import (
	"testing"

	"github.com/hlmerscher/go-algo/misc"
)

func TestMergeSort(t *testing.T) {
	t.Run("Test Simple Sorting", misc.TestSimpleSorting(MergeSort))
	t.Run("Test Random Collection Sorting", misc.TestSortingRandomCollection(MergeSort))
}

func BenchmarkMergeSortWorstCase(b *testing.B) {
	col := misc.GenInvertedCollectionWithDefault()
	MergeSort(col)
}
