package datastructures

import "fmt"

// TrieNode represents each node in the trie
type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

// Trie represents the trie itself
type Trie struct {
	root *TrieNode
}

// NewTrie creates a new trie
func NewTrie() *Trie {
	return &Trie{root: &TrieNode{children: make(map[rune]*TrieNode)}}
}

// Insert adds a word to the trie
func (t *Trie) Insert(word string) {
	node := t.root
	for _, char := range word {
		if node.children[char] == nil {
			node.children[char] = &TrieNode{children: make(map[rune]*TrieNode)}
		}
		node = node.children[char]
	}
	node.isEnd = true
}

// Search checks if a word is in the trie
func (t *Trie) Search(word string) bool {
	node := t.root
	for _, char := range word {
		if node.children[char] == nil {
			return false
		}
		node = node.children[char]
	}
	return node.isEnd
}

func MainTries() {
	trie := NewTrie()

	// Insert words into the trie
	trie.Insert("cat")
	trie.Insert("car")
	trie.Insert("dog")

	// Search for words in the trie
	fmt.Println(trie.Search("cat")) // true
	fmt.Println(trie.Search("car")) // true
	fmt.Println(trie.Search("dog")) // true
	fmt.Println(trie.Search("cow")) // false
}
