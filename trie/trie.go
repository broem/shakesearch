package trie

import "pulley.com/shakesearch/models"

type Trie struct {
	root *models.TrieNode
}

func NewTrie() *Trie {
	return &Trie{root: &models.TrieNode{Children: make(map[rune]*models.TrieNode)}}
}

func (t *Trie) Insert(word string) {
	currentNode := t.root
	for _, char := range word {
		if currentNode.Children[char] == nil {
			currentNode.Children[char] = &models.TrieNode{Children: make(map[rune]*models.TrieNode)}
		}
		currentNode = currentNode.Children[char]
	}
	currentNode.EndOfWord = true
}

func (t *Trie) Search(prefix string) []string {
	currentNode := t.root
	for _, char := range prefix {
		if currentNode.Children[char] == nil {
			return []string{}
		}
		currentNode = currentNode.Children[char]
	}
	return t.findWordsFrom(currentNode, prefix, []string{})
}

func (t *Trie) findWordsFrom(node *models.TrieNode, prefix string, words []string) []string {
	if node.EndOfWord {
		words = append(words, prefix)
	}
	for char, childNode := range node.Children {
		words = t.findWordsFrom(childNode, prefix+string(char), words)
	}
	return words
}

func (t *Trie) FuzzySearch(target string, maxCost int) []string {
	currentRow := make([]int, len(target)+1)

	for i := range currentRow {
		currentRow[i] = i
	}

	results := []string{}

	// Recursively search each branch of the Trie
	for char, node := range t.root.Children {
		t.searchRecursive(node, &results, target, string(char), currentRow, maxCost)
	}

	return results
}

func (t *Trie) searchRecursive(node *models.TrieNode, results *[]string, target string, currentWord string,
	currentRow []int, maxCost int) {

	columns := len(target) + 1
	newRow := make([]int, columns)
	newRow[0] = len(currentWord)

	// Fill up the newRow array
	minCost := int(^uint(0) >> 1) // Max int
	for col := 1; col < columns; col++ {
		insertionCost := currentRow[col] + 1
		deletionCost := newRow[col-1] + 1
		var replacementCost int
		if target[col-1] != currentWord[len(currentWord)-1] {
			replacementCost = currentRow[col-1] + 1
		} else {
			replacementCost = currentRow[col-1]
		}

		newRow[col] = min(insertionCost, deletionCost, replacementCost)

		if newRow[col] < minCost {
			minCost = newRow[col]
		}
	}

	if newRow[len(newRow)-1] <= maxCost && node.EndOfWord {
		*results = append(*results, currentWord)
	}

	if minCost <= maxCost {
		for char, childNode := range node.Children {
			t.searchRecursive(childNode, results, target, currentWord+string(char), newRow, maxCost)
		}
	}
}

func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	}
	if b < c {
		return b
	}
	return c
}