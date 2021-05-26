package sorting

import (
	"testing"

	"github.com/hlmerscher/go-algo/misc"
)

func TestBubbleSort(t *testing.T) {
	t.Run("Test Simple Sorting", misc.TestSimpleSorting(BubbleSort))
	t.Run("Test Random Collection Sorting", misc.TestSortingRandomCollection(BubbleSort))
}

func BenchmarkBubbleSortWorstCase(b *testing.B) {
	col := misc.GenInvertedCollectionWithDefault()
	BubbleSort(col)
}
