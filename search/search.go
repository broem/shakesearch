package search

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"pulley.com/shakesearch/indexer"
	"pulley.com/shakesearch/models"
	"pulley.com/shakesearch/trie"
)

type Searcher struct {
	Index     map[string][]models.IndexEntry
	Documents map[string]string
	Trie      *trie.Trie
}

func (s *Searcher) SearchKeyword(keyword string) {
	entries, found := s.Index[strings.ToLower(keyword)]
	if !found {
		fmt.Printf("Keyword '%s' not found in any document.\n", keyword)
		return
	}

	fmt.Printf("Keyword '%s' found in:\n", keyword)
	for _, entry := range entries {
		fmt.Printf("Document: '%s', Line: %d\n", entry.DocumentID, entry.Position)
	}
}

func (s *Searcher) SearchKeywordWithContext( keyword string, contextLen int, documentIDs ...string) []models.SearchMatch {
	entries, found := s.Index[strings.ToLower(keyword)]
	if !found {
		fmt.Printf("Keyword '%s' not found in any document.\n", keyword)
		return nil
	}

	// Convert the list of document IDs to a map for faster lookup
	documentIDMap := make(map[string]bool)
	for _, id := range documentIDs {
		documentIDMap[id] = true
	}

	matches := make([]models.SearchMatch, 0)
	for _, entry := range entries {
		// If document IDs were provided, and the current document ID is not in the list, skip it
		if len(documentIDMap) > 0 && !documentIDMap[entry.DocumentID] {
			continue
		}

		text := s.Documents[entry.DocumentID]

		start := entry.Position - contextLen
		if start < 0 {
			start = 0
		}
		end := entry.Position + len(keyword) + contextLen
		if end > len(text) {
			end = len(text)
		}
		context := text[start:end]

		match := models.SearchMatch{DocumentID: entry.DocumentID, Position: entry.Position, Context: context}
		matches = append(matches, match)
	}

	return matches
}

func (s *Searcher) Load(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return err
	}
	defer file.Close()

	var documentID, previousLine string
	var lines []string
	documents := make(map[string]string)
	firstContent := true
	firstSonnets := true

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		switch {
		case strings.Contains(line, "THE SONNETS"):
			if firstSonnets {
				firstSonnets = false
				continue
			}
			if documentID != "" {
				documents[documentID] = strings.Join(lines, "\n")
				lines = nil
			}
			documentID = "THE SONNETS"
		case strings.Contains(line, "Contents"):
			if firstContent {
				firstContent = false
				continue
			}
			if documentID != "" {
				documents[documentID] = strings.Join(lines, "\n")
				lines = nil
			}
			documentID = strings.TrimSpace(previousLine)
		case strings.Contains(line, "DRAMATIS PERSONAE"):
			if documentID != "" {
				documents[documentID] = strings.Join(lines, "\n")
				lines = nil
			}
			documentID = strings.TrimSpace(previousLine)
		case documentID != "" && (strings.Contains(line, "THE END") || strings.Contains(line, "Contents") || strings.Contains(line, "DRAMATIS PERSONAE")):
			documents[documentID] = strings.Join(lines, "\n")
			lines = nil
			documentID = ""
		default:
			if documentID != "" {
				lines = append(lines, line)
			}
		}

		if line != "" {
			previousLine = line
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
		return err
	}

	// Now documents map contains all the parsed sections.
	// You can turn each section into an inverted index.
	index := indexer.CreateInvertedIndex(documents)

	trie := trie.NewTrie()
	for word := range index {
		trie.Insert(word)
	}

	s.Index = index
	s.Documents = documents
	s.Trie = trie

	return nil
}

func (s *Searcher) Search(query string, documents... string) []models.SearchMatch {
	results := s.SearchKeywordWithContext(strings.ToLower(query), 250, documents...)
	return results
}
