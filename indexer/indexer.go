package indexer

import (
	"bufio"
	"strings"
	"unicode"

	"pulley.com/shakesearch/models"
)

func CreateInvertedIndex(documents map[string]string) map[string][]models.IndexEntry {
	invertedIndex := make(map[string][]models.IndexEntry)

	for documentID, text := range documents {
		scanner := bufio.NewScanner(strings.NewReader(text))
		scanner.Split(bufio.ScanRunes)

		position := 0
		word := ""
		for scanner.Scan() {
			char := scanner.Text()

			if unicode.IsSpace(rune(char[0])) || unicode.IsPunct(rune(char[0])) {
				if len(word) > 0 {
					wordLower := strings.ToLower(word)
					invertedIndex[wordLower] = append(invertedIndex[wordLower], models.IndexEntry{DocumentID: documentID, Position: position - len(word)})
					word = ""
				}
			} else {
				word += char
			}

			position += len(char)
		}
	}

	return invertedIndex
}

func CreateAutocompleteIndex(invertedIndex map[string][]models.IndexEntry) map[string][]models.IndexEntry {
	autocompleteIndex := make(map[string][]models.IndexEntry)

	for word, entries := range invertedIndex {
		for i := 1; i <= len(word); i++ {
			prefix := word[:i]
			autocompleteIndex[prefix] = append(autocompleteIndex[prefix], entries...)
		}
	}

	return autocompleteIndex
}