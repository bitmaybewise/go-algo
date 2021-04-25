package algo

import (
	"testing"
)

func TestIntSet(t *testing.T) {
	set := NewSet()
	set.Add(1)
	set.Add(1)
	set.Add(2)
	set.Add(2)
	set.Add(2)
	set.Add(3)

	if set.Size() > 3 {
		t.Error("Set has duplicated items")
	}

	items := []int{1, 2, 3}
	if set.Contains(items...) {
		t.Errorf("Set does not contain items %v, but %v\n", items, set.Values())
	}
}

func TestSetFromValues(t *testing.T) {
	values := []int{1, 1, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3}
	set := SetFromValues(values...)

	if set.Size() > 3 {
		t.Error("Set has duplicated items")
	}

	items := []int{1, 2, 3}
	if set.Contains(items...) {
		t.Errorf("Set does not contain items %v, but %v\n", items, set.Values())
	}
}
