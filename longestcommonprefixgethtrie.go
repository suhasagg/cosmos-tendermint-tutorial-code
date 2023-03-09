package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/trie"
)

func main() {
	strs := []string{"apple", "apricot", "apology", "apartment", "apostle", "applepie", "apartmentcomplex"}

	// Create new trie
	t := trie.NewTrie()

	// Insert all strings into the trie
	for _, s := range strs {
		key := []byte(s)
		t.Update(key, key)
	}

	// Get longest common prefix
	prefix, _, _ := longestCommonPrefix(t)
	fmt.Println(string(prefix))
}

func longestCommonPrefix(t *trie.Trie) ([]byte, bool, error) {
	// Get root node
	root := t.Node()

	// Initialize prefix
	var prefix []byte

	// Traverse trie until a leaf node is reached or multiple children are found
	for {
		// Check if node is a leaf
		if len(root.Key) != 0 {
			return prefix, true, nil
		}

		// Check number of children
		if len(root.Children) != 1 {
			return prefix, false, nil
		}

		// Get child node
		child := root.Children[0]

		// Add child's prefix to prefix
		prefix = append(prefix, child.Key...)

		// Set root to child
		root = child
	}
}
