package algo

import (
	"fmt"
	"testing"
)

func TestGeneratePermutations(t *testing.T) {
	for i := 0; i < 4; i++ {
		t.Run(fmt.Sprintf("Permute %d", i), func(t *testing.T) {
			p := GeneratePermutations(i)

			if fac := Factorial(i); len(p) != fac {
				t.Errorf("Should have %d permutations instead of %d", fac, len(p))
			}

			resultsAsStrings := make(map[string]int)
			for _, result := range p {
				resultsAsStrings[fmt.Sprintf("%v", result)]++
				set := SetFromValues(result...)

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
