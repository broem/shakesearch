package handlers

import (
	"net/http"
	"strings"

	"pulley.com/shakesearch/search"
)

func HandleSuggest(searcher search.Searcher) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		query, ok := r.URL.Query()["q"]
		if !ok || len(query[0]) < 1 {
			writeResponseWithStatusCode(w, "missing search query in URL params", http.StatusBadRequest)
			return
		}

		lowerCaseQuery := strings.ToLower(query[0])
		suggestions := searcher.Trie.Search(lowerCaseQuery)

		// limit the number of suggestions to 5
		if len(suggestions) > 5 {
			suggestions = suggestions[:5]
		}

		writeJSONResponse(w, suggestions)
	}
}