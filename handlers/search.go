package handlers

import (
	"bytes"
	"encoding/json"
	"math"
	"net/http"
	"strconv"
	"strings"

	"pulley.com/shakesearch/models"
	"pulley.com/shakesearch/search"
)

func HandleSearch(searcher search.Searcher) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		query, ok := r.URL.Query()["q"]
		if !ok || len(query[0]) < 1 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("missing search query in URL params"))
			return
		}

		page, _ := strconv.Atoi(r.URL.Query().Get("page"))
		if page < 1 {
			page = 1
		}
		pageSize, _ := strconv.Atoi(r.URL.Query().Get("pageSize"))
		if pageSize < 1 {
			pageSize = 10
		}

		documents := r.URL.Query()["works"]
		if len(documents) == 1 {
			documents = append(strings.Split(documents[0], ","), documents[1:]...)
		}

		allResults := searcher.Search(query[0], documents...)
		totalResults := len(allResults)
		totalPages := int(math.Ceil(float64(totalResults) / float64(pageSize)))

		// Now apply pagination
		start := (page - 1) * pageSize
		end := start + pageSize
		if end > totalResults {
			end = totalResults
		}

		// need to make sure start and end are not within the same page
		if start > end {
			writeResponseWithStatusCode(w, "invalid page or pageSize", http.StatusBadRequest)
			return
		}

		results := allResults[start:end]

		// Create a response struct to include results and pagination info
		response := models.SearchResponse{
			Results:      results,
			TotalResults: totalResults,
			Page:         page,
			TotalPages:   totalPages,
		}
		buf := &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		err := enc.Encode(response)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("encoding failure"))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(buf.Bytes())
	}
}