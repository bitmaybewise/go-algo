package algo

import (
	"math"
)

// Knapsack data structure to hold the values that fit on the knapsack
type Knapsack struct {
	Amount int
}

// KnapsackUnit data structure to represent items of a Knapsack
type KnapsackUnit struct {
	Weight int
	Value  int
}

// NewKnapsack implementation of a knapsack algorithm using dynamic programming
func NewKnapsack(totalCapacity int, units []KnapsackUnit) *Knapsack {
	acc := make([][]int, len(units)+1)
	for i := range acc {
		acc[i] = make([]int, totalCapacity+1)
	}

	for i := 1; i <= len(units); i++ {
		for j := 1; j <= totalCapacity; j++ {
			if j >= units[i-1].Weight {
				v1 := float64(acc[i-1][j])
				v2 := float64(units[i-1].Value + acc[i-1][j-units[i-1].Weight])
				acc[i][j] = int(math.Max(v1, v2))
			} else {
				acc[i][j] = acc[i-1][j]
			}
		}
	}

	return &Knapsack{acc[len(units)][totalCapacity]}
}
