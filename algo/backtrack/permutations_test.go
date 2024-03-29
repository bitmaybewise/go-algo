package backtrack

import (
	"fmt"
	"testing"

	"github.com/hlmerscher/go-algo/algo"
	"github.com/hlmerscher/go-algo/misc"
)

func TestGeneratePermutations(t *testing.T) {
	for i := 0; i < 4; i++ {
		t.Run(fmt.Sprintf("Permute %d", i), func(t *testing.T) {
			p := GeneratePermutations(i)

			if fac := misc.Factorial(i); len(p) != fac {
				t.Errorf("Should have %d permutations instead of %d", fac, len(p))
			}

			resultsAsStrings := make(map[string]int)
			for _, result := range p {
				resultsAsStrings[fmt.Sprintf("%v", result)]++
				set := algo.SetFromValues(result...)

				if set.Size() != i {
					t.Errorf("Permutation has duplicates: %v\n", result)
				}
			}

			for result, count := range resultsAsStrings {
				if count > 1 {
					t.Errorf("Combination %v is duplicated", result)
				}
			}
		})
	}
}

func TestGeneratePermutationsWithList(t *testing.T) {
	inputs := [][]int{
		{666},
		{3, 5},
		{3, 5, 7},
	}
	for _, col := range inputs {
		t.Run(fmt.Sprintf("Permute %v", col), func(t *testing.T) {
			p := GeneratePermutationsWithList(col)

			fac := misc.Factorial(len(col))
			if len(p) != fac {
				t.Errorf("Should have %d permutations instead of %d", fac, len(p))
			}

			resultsAsStrings := make(map[string]int)
			for _, result := range p {
				resultsAsStrings[fmt.Sprintf("%v", result)]++
				set := algo.SetFromValues(result...)

				if set.Size() != len(col) {
					t.Errorf("Permutation has duplicates: %v\n", result)
				}
			}

			for result, count := range resultsAsStrings {
				if count > 1 {
					t.Errorf("Combination %v is duplicated", result)
				}
			}
		})
	}
}
