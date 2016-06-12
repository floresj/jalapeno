package web

import (
	"encoding/json"
	"net/http"
)

// Helper that json encodes any object and writes it to a response writer
func EncodeJson(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(data)
}

// Helper that json decodes a request body to the specified interface
func DecodeJson(r *http.Request, v interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return err
	}
	return nil
}
