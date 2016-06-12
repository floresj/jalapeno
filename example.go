package main

import (
	"fmt"
	"net/http"

	"github.com/floresj/jalapeno/web"
)

func main() {
	m := http.NewServeMux()

	// Create new chain with Err handler and any number of middleware funcs
	c := web.NewChain(ErrorHandler, Logger)

	// For your individual handlers, pass in an Endpoint
	m.HandleFunc("/users", c.Endpoint(UserHandler))
	m.HandleFunc("/", c.Endpoint(HomeHandler))

	http.ListenAndServe(":8888", m)
}

func UserHandler(w http.ResponseWriter, r *http.Request) error {
	return web.EncodeJson(w, struct {
		Firstname string
		Lastname  string
	}{
		Firstname: "John",
		Lastname:  "Flores",
	})
}

func HomeHandler(w http.ResponseWriter, r *http.Request) error {
	fmt.Fprintf(w, "Hey man")
	return nil
}

func ErrorHandler(f web.EndpointFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := f(w, r)
		if err == nil {
			return
		}
		switch err {
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func Logger(f web.EndpointFunc) web.EndpointFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		fmt.Println("I like to print stuff out")
		return f(w, r)
	}
}
