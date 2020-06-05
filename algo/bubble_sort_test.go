package algo

import (
	"testing"

	"github.com/hlmerscher/go-algo/helper"
)

func TestBubbleSort(t *testing.T) {
	t.Run("Test Simple Sorting", helper.TestSimpleSorting(BubbleSort))
	t.Run("Test Random Collection Sorting", helper.TestSortingRandomCollection(BubbleSort))
}

func BenchmarkBubbleSortWorstCase(b *testing.B) {
	col := helper.GenInvertedCollectionWithDefault()
	BubbleSort(col)
}