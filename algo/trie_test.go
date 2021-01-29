package algo

import (
	"testing"
)

func TestTrieAutocomplete(t *testing.T) {
	t.Run("Without entries", func(t *testing.T) {
		trie := NewTrie()

		if result := trie.Autocomplete("ca"); len(result) != 0 {
			t.Errorf("Result should be empty, instead returned %v", result)
		}
	})
	t.Run("With entries", func(t *testing.T) {
		trie := NewTrie()
		trie.Insert("cab")
		trie.Insert("cat")
		trie.Insert("tab")

		result := trie.Autocomplete("ca")
		if len(result) != 2 {
			t.Errorf("Result should contain %d, instead got %d", 2, len(result))
		}
		items := []string{"cab", "cat"}
		for i, name := range items {
			if result[i] != name {
				t.Errorf("Index %c should have value %v, instead got %v", i, name, result[i])
			}
		}
	})
}
