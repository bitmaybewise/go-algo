package algo

import (
	"testing"

	"github.com/hlmerscher/go-algo/misc"
)

func TestKthLowest(t *testing.T) {
	collection := []int{1, 50, 20, 10, 60, 30}

	t.Run("First lowest", func(t *testing.T) {
		if value := KthLowest(collection, 0); value != 1 {
			t.Errorf("Should return %d but instead returned %d\n", 1, value)
		}
	})
	t.Run("Second lowest", func(t *testing.T) {
		if value := KthLowest(collection, 1); value != 10 {
			t.Errorf("Should return %d but instead returned %d\n", 10, value)
		}
	})
	t.Run("Last", func(t *testing.T) {
		if value := KthLowest(collection, len(collection)-1); value != 60 {
			t.Errorf("Should return %d but instead returned %d\n", 60, value)
		}
	})
}

func BenchmarkKthLowest(b *testing.B) {
	collection := misc.GenRandomCollection(misc.Limit)
	KthLowest(collection, 0)
}
