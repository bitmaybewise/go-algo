package helper

import (
	"math/rand"
	"testing"
)

const Limit = 10000

// GenRandomCollection generates random collection of integers up to N size
func GenRandomCollection(n int) []int {
	col := make([]int, n)
	for i := 0; i < n; i++ {
		col[i] = rand.Intn(n)
	}
	return col
}

// GenOrderedCollectionWithDefault generates a collection with values in ascent order
func GenOrderedCollectionWithDefault() []int {
	return GenOrderedCollection(Limit)
}

// GenOrderedCollection generates a collection with values in ascent order
func GenOrderedCollection(n int) []int {
	col := make([]int, n)
	for i := 0; i < n; i++ {
		col[i] = i
	}
	return col
}

// GenInvertedCollectionWithDefault generates a collection with values in descent order
func GenInvertedCollectionWithDefault() []int {
	return GenInvertedCollection(Limit)
}

// GenInvertedCollection generates a collection with values in descent order
func GenInvertedCollection(n int) []int {
	col := make([]int, n)
	for i := 0; i < n; i++ {
		col[i] = n - i
	}
	return col
}

// IsSorted checks if collection of integers is sorted
func IsSorted(col []int) bool {
	size := len(col)
	for i := 0; i < size-1; i++ {
		nextIdx := i + 1
		if col[i] > col[nextIdx] {
			return false
		}
	}
	return true
}

// TestSimpleSorting DRY function to test a very simple sorting case
func TestSimpleSorting(fn func(col []int)) func(t *testing.T) {
	return func(t *testing.T) {
		col := []int{7, 5, 1, 8, 3}
		expected := []int{1, 3, 5, 7, 8}
		fn(col)

		for i := 0; i < len(col); i++ {
			if col[i] != expected[i] {
				t.Errorf("Value %d at pos %d should be %d", col[i], i, expected[i])
			}
		}
	}
}

// TestSortingRandomCollection DRY function to test a random collection sorting case
func TestSortingRandomCollection(fn func(col []int)) func(t *testing.T) {
	return func(t *testing.T) {
		col := GenRandomCollection(Limit)

		if fn(col); !IsSorted(col) {
			t.Errorf("Collection is not sorted\n%v", col)
		}
	}
}
