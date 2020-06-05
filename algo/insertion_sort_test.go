package algo

import (
	"testing"

	"github.com/hlmerscher/go-algo/helper"
)

func TestInsertionSort(t *testing.T) {
	t.Run("Test Simple Sorting", helper.TestSimpleSorting(InsertionSort))
	t.Run("Test Random Collection Sorting", helper.TestSortingRandomCollection(InsertionSort))
}

func BenchmarkInsertionSortWorstCase(b *testing.B) {
	col := helper.GenInvertedCollectionWithDefault()
	InsertionSort(col)
}
