package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func writeError(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func writeResponseWithStatusCode(w http.ResponseWriter, data interface{}, statusCode int) {
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	err := enc.Encode(data)
	if err != nil {
		writeError(w, err)
		return
	}
	w.WriteHeader(statusCode)
	w.Write(buf.Bytes())
}

func writeJSONResponse(w http.ResponseWriter, data interface{}) {
    buf := &bytes.Buffer{}
    enc := json.NewEncoder(buf)
    err := enc.Encode(data)
    if err != nil {
        writeError(w, err)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.Write(buf.Bytes())
}