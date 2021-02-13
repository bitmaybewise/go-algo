package algo

type TrieNodeChildren map[rune]*TrieNode

type TrieNode struct {
	children TrieNodeChildren
}

type Trie struct {
	root *TrieNode
}

func newTrieNode() TrieNode {
	return TrieNode{children: make(TrieNodeChildren)}
}

func NewTrie() Trie {
	node := newTrieNode()
	return Trie{root: &node}
}

func (t *Trie) Insert(word string) {
	node := t.root

	for _, char := range word {
		if child, ok := node.children[char]; ok {
			node = child
		} else {
			newNode := newTrieNode()
			node.children[char] = &newNode
			node = &newNode
		}
	}

	node.children['*'] = nil
}

func (t *Trie) search(word string) *TrieNode {
	node := t.root
	if node == nil {
		return nil
	}

	for _, char := range word {
		if _, ok := node.children[char]; ok {
			node = node.children[char]
		} else {
			return nil
		}
	}

	return node
}

func (t *Trie) collectAllWords(node *TrieNode, word string, words []string) []string {
	for char, child := range node.children {
		if char == '*' {
			words = append(words, word)
		} else {
			words = t.collectAllWords(child, word+string(char), words)
		}
	}
	return words
}

func (t *Trie) Autocomplete(prefix string) []string {
	node := t.search(prefix)
	if node == nil {
		return nil
	}
	words := []string{}
	return t.collectAllWords(node, prefix, words)
}
