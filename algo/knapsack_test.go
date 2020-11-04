package algo

import (
	"testing"
)

func TestKnapsack(t *testing.T) {
	cap := 4
	guitar := KnapsackUnit{Weight: 1, Value: 1500}
	stereo := KnapsackUnit{Weight: 4, Value: 3000}
	laptop := KnapsackUnit{Weight: 3, Value: 2000}
	iphone := KnapsackUnit{Weight: 1, Value: 2000}

	t.Run("Without units", func(t *testing.T) {
		k := NewKnapsack(cap, make([]KnapsackUnit, 0))

		if k.Amount != 0 {
			t.Errorf("Amount should be %d, got %d", 0, k.Amount)
		}
	})

	t.Run("With a single unit", func(t *testing.T) {
		k := NewKnapsack(cap, []KnapsackUnit{guitar})

		if k.Amount != guitar.Value {
			t.Errorf("Amount should be %d, got %d", guitar.Value, k.Amount)
		}
	})

	t.Run("With 2 units", func(t *testing.T) {
		k := NewKnapsack(cap, []KnapsackUnit{guitar, stereo})
		amount := stereo.Value

		if k.Amount != amount {
			t.Errorf("Amount should be %d, got %d", amount, k.Amount)
		}
	})

	t.Run("With 3 units", func(t *testing.T) {
		k := NewKnapsack(cap, []KnapsackUnit{guitar, stereo, laptop})
		amount := guitar.Value + laptop.Value

		if k.Amount != amount {
			t.Errorf("Amount should be %d, got %d", amount, k.Amount)
		}
	})

	t.Run("With 4 units", func(t *testing.T) {
		k := NewKnapsack(cap, []KnapsackUnit{guitar, stereo, laptop, iphone})
		amount := laptop.Value + iphone.Value

		if k.Amount != amount {
			t.Errorf("Amount should be %d, got %d", amount, k.Amount)
		}
	})
}
