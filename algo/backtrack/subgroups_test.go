package backtrack

import (
	"fmt"
	"testing"
)

func TestGenerateSubgroups(t *testing.T) {
	acc := make(map[string]bool)
	for _, subgroup := range GenerateSubgroups(3) {
		acc[fmt.Sprintf("%v", subgroup)] = true
	}

	expectedSubgroups := []string{
		"[]",
		"[1]",
		"[2]",
		"[3]",
		"[1 2]",
		"[1 3]",
		"[2 3]",
		"[1 2 3]",
	}
	for _, value := range expectedSubgroups {
		if ok := acc[value]; !ok {
			t.Errorf("Subgroup %s not found", value)
		}
	}
}
