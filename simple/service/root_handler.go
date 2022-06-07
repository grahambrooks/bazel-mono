package main

import (
	"encoding/json"
	"net/http"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var body = map[string]interface{}{
		"service": "Simple Service",
	}

	encoder := json.NewEncoder(w)

	_ = encoder.Encode(body)
}
