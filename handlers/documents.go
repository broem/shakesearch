package handlers

import (
	"net/http"
	"sort"

	"pulley.com/shakesearch/search"
)

func HandleDocuments(searcher search.Searcher) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// pull the document keys out of the searcher
		var documentKeys []string
		for key := range searcher.Documents {
			documentKeys = append(documentKeys, key)
		}

		// sort the document keys
		sort.Strings(documentKeys)

		writeJSONResponse(w, documentKeys)
	}
}