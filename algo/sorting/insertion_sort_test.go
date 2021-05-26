package sorting

import (
	"testing"

	"github.com/hlmerscher/go-algo/misc"
)

func TestInsertionSort(t *testing.T) {
	t.Run("Test Simple Sorting", misc.TestSimpleSorting(InsertionSort))
	t.Run("Test Random Collection Sorting", misc.TestSortingRandomCollection(InsertionSort))
}

func BenchmarkInsertionSortWorstCase(b *testing.B) {
	col := misc.GenInvertedCollectionWithDefault()
	InsertionSort(col)
}
