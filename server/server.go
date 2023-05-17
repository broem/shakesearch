package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"pulley.com/shakesearch/handlers"
	"pulley.com/shakesearch/search"
)

func Start() {
	searcher := search.Searcher{}
	searcher.Load("completeworks.txt")

	fs := http.FileServer(http.Dir("./shakesearch/build"))

	http.Handle("/", fs)

	http.Handle("/search", corsMiddleware(http.HandlerFunc(handlers.HandleSearch(searcher))))
	http.Handle("/suggest", corsMiddleware(http.HandlerFunc(handlers.HandleSuggest(searcher))))
	http.Handle("/documents", corsMiddleware(http.HandlerFunc(handlers.HandleDocuments(searcher))))

	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}

	fmt.Printf("Listening on port %s...", port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		log.Fatal(err)
	}
}

func corsMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        enableCors(&w)
        next.ServeHTTP(w, r)
    })
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
