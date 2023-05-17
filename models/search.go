package models

type SearchResponse struct {
	Page         int           `json:"page"`
	Results      []SearchMatch `json:"results"`
	TotalPages   int           `json:"totalPages"`
	TotalResults int           `json:"totalResults"`
}

type SearchMatch struct {
	DocumentID string `json:"documentId"`
	Position   int    `json:"position"`
	Context    string `json:"context"`
}

type DocumentInfo struct {
	DocumentID string
	Position   int
}

type InvertedIndex map[string][]DocumentInfo

type SearchResult struct {
	Found   bool
	Matches []SearchMatch
	Message string
}