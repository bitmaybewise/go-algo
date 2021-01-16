package algo

import (
	"testing"
)

func TestMaxHeap(t *testing.T) {
	t.Run("Insert", func(t *testing.T) {
		heap := NewMaxHeap()
		heap.Insert(10)
		heap.Insert(2)
		heap.Insert(20)
		heap.Insert(11)
		heap.Insert(9)
		expectedOrder := []int{20, 11, 10, 2, 9}

		if len(heap.data) != len(expectedOrder) {
			t.Errorf("Expected value %v, result %v", expectedOrder, heap.data)
			return
		}

		for i := range heap.data {
			if heap.data[i] != expectedOrder[i] {
				t.Errorf("Expected value %v, result %v", expectedOrder, heap.data)
				break
			}
		}
	})
	t.Run("Pop", func(t *testing.T) {
		heap := NewMaxHeap()
		heap.Insert(10)
		heap.Insert(2)
		heap.Insert(20)
		heap.Insert(11)
		heap.Insert(9)

		value := heap.Pop()
		if value != 20 {
			t.Errorf("Root deleted should be %d, got %d", 20, value)
		}
		value = heap.Pop()
		if value != 11 {
			t.Errorf("Root deleted should be %d, got %d", 11, value)
		}

		expectedOrder := []int{10, 9, 2}
		if len(heap.data) != len(expectedOrder) {
			t.Errorf("Expected value %v, result %v", expectedOrder, heap.data)
			return
		}

		for i := range heap.data {
			if heap.data[i] != expectedOrder[i] {
				t.Errorf("Expected value %v, result %v", expectedOrder, heap.data)
				break
			}
		}
	})
	t.Run("Root", func(t *testing.T) {
		heap := NewMaxHeap()

		heap.Insert(10)
		value := heap.Root()
		if value != 10 {
			t.Errorf("Root should be %d, got %d", 10, value)
		}

		heap.Insert(2)
		value = heap.Root()
		if value != 10 {
			t.Errorf("Root should be %d, got %d", 10, value)
		}

		heap.Insert(20)
		value = heap.Root()
		if value != 20 {
			t.Errorf("Root should be %d, got %d", 10, value)
		}
	})
	t.Run("Last", func(t *testing.T) {
		heap := NewMaxHeap()

		heap.Insert(10)
		value := heap.Last()
		if value != 10 {
			t.Errorf("Last should be %d, got %d", 10, value)
		}

		heap.Insert(2)
		value = heap.Last()
		if value != 2 {
			t.Errorf("Last should be %d, got %d", 2, value)
		}

		heap.Insert(20)
		value = heap.Last()
		if value != 10 {
			t.Errorf("Last should be %d, got %d", 10, value)
		}
	})
}
