package algo

import (
	"testing"

	"github.com/hlmerscher/go-algo/helper"
)

func TestBinarySearch(t *testing.T) {
	collection := helper.GenOrderedCollectionWithDefault()

	t.Run("Not found with single item", func(t *testing.T) {
		if idx := BinarySearch([]int{7}, -5); idx != -1 {
			t.Errorf("Should return -1 but instead returned %d\n", idx)
		}
	})
	t.Run("Not found with multiple items", func(t *testing.T) {
		if idx := BinarySearch(collection, -5); idx != -1 {
			t.Errorf("Should return -1 but instead returned %d\n", idx)
		}
	})
	t.Run("Found with single item", func(t *testing.T) {
		if idx := BinarySearch([]int{7}, 7); idx != 0 {
			t.Errorf("Should return %d but instead returned %d\n", 7, idx)
		}
	})
	t.Run("Found with multiple items", func(t *testing.T) {
		if idx := BinarySearch(collection, collection[5000]); idx != 5000 {
			t.Errorf("Should return %d but instead returned %d\n", collection[5000], idx)
		}
	})
}

func BenchmarkBinarySearch(b *testing.B) {
	collection := helper.GenOrderedCollectionWithDefault()
	BinarySearch(collection, 1)
}
