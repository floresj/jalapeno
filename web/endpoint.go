package web

import "net/http"

// EndpointFunc is a function with access to Request and Response values that returns an error.
// This is meant to be used as an http handler
type EndpointFunc func(w http.ResponseWriter, r *http.Request) error

// EndpointFuncChain wraps an EndpointFunc with another Endpoint func so that they combine powers
type EndpointFuncChain func(EndpointFunc) EndpointFunc

// EndpointFuncErrHandler is a custom handler that that takes in an EndpointFunc and returns
// a http.HandlerFunc
type EndpointFuncErrHandler func(EndpointFunc) http.HandlerFunc
