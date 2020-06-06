package algo

import (
	"testing"

	"github.com/hlmerscher/go-algo/helper"
)

func TestMergeSort(t *testing.T) {
	t.Run("Test Simple Sorting", helper.TestSimpleSorting(MergeSort))
	t.Run("Test Random Collection Sorting", helper.TestSortingRandomCollection(MergeSort))
}

func BenchmarkMergeSortWorstCase(b *testing.B) {
	col := helper.GenInvertedCollectionWithDefault()
	MergeSort(col)
}
