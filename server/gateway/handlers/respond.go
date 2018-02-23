package handlers

import (
"net/http"
"encoding/json"
"fmt"
)

//respond encodes `value` into JSON and writes that to the response
func respond(w http.ResponseWriter, value interface{}) {
	w.Header().Add(headerContentType, contentTypeJSON)
	if err := json.NewEncoder(w).Encode(value); err != nil {
		http.Error(w, fmt.Sprintf("error encoding response value to JSON: %v", err), http.StatusInternalServerError)
	}
}
