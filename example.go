package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/floresj/jalapeno/web"
)

// ErrForStuff is a sample error just to demonstrate possible usage of ErrorHandler
var ErrForStuff = errors.New("Some stuff went wrong")

func main() {
	m := http.NewServeMux()

	// Create new chain with Err handler and any number of middleware funcs
	c := web.NewChain(ErrorHandler, Auth, Logger)

	// For your individual handlers, pass in an Endpoint
	m.HandleFunc("/users", c.Endpoint(UserHandler))
	m.HandleFunc("/", c.Endpoint(HomeHandler))

	http.ListenAndServe(":8888", m)
}

func UserHandler(w http.ResponseWriter, r *http.Request) error {
	return web.EncodeJSON(w, struct {
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

func StuffHandler(w http.ResponseWriter, r *http.Request) error {
	// Lets assume some stuff happened and it went all wrong
	return ErrForStuff
}

func ErrorHandler(f web.EndpointFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := f(w, r)
		if err == nil {
			return
		}
		switch err {
		case ErrForStuff:
			http.Error(w, err.Error(), http.StatusNotFound)
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

func Auth(f web.EndpointFunc) web.EndpointFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		fmt.Println("Maybe you want to check credentials before moving on to the next handler")
		return f(w, r)
	}
}
