package web

import "net/http"

// EndpointFunc is a function that returns an error.
type EndpointFunc func(w http.ResponseWriter, r *http.Request) error
