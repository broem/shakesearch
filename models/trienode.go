package models

type TrieNode struct {
	Children  map[rune]*TrieNode
	EndOfWord bool
}
