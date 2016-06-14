package web

import (
	"net/http"
)

// Chain enables EndpointFunc's to be defined and executed like middleware.
type Chain struct {
	// Manages errors returned by EndpointFunc's. Executes the passed in EndpointFunc
	// and then determines which error to return. If nil, simply returns. If non-nil,
	// checks against a pre-defined set of errors and determines its corresponding status
	// and error message and writes it to the ResponseWriter.
	//
	// This alleviates the need for EndpointFuncs to worry about handling error messages. Instead,
	// EndpointFunc's simply return nil or an error and this function properly manages those errors.
	errHandler EndpointFuncErrHandler

	funcs []EndpointFuncChain
}

// NewChain creates a Chain of handlers to be executed in the order that they are defined
// Example:
//	NewChain(Func1, Func2, Func3)
//
func NewChain(errHandler EndpointFuncErrHandler, funcs ...EndpointFuncChain) Chain {
	c := Chain{}
	if errHandler != nil {
		c.errHandler = errHandler
	} else {
		c.errHandler = defaultErrHandler
	}
	// Appends the variant endpoint funcs to the slice of EndpointFuncs
	c.funcs = append(c.funcs, funcs...)
	return c
}

// Endpoint sets up the http.HandlerFunc along with the middleware functions.
// Example:
//	NewChain(Func1, Func2, Func3).Endpoint(Func4)
//	This would be executed as follows:
//	Func1(Func2(Func3(Func4(h))))
// If a func returns an error, the call cycle is short circuited.
func (c Chain) Endpoint(endpoint EndpointFunc) http.HandlerFunc {
	var final EndpointFunc
	if endpoint != nil {
		final = endpoint
	}

	for i := len(c.funcs) - 1; i >= 0; i-- {
		final = c.funcs[i](final)
	}

	return c.errHandler(func(w http.ResponseWriter, r *http.Request) error {
		return final(w, r)
	})
}

// defaultErrHandler is the Default error handler used if one is not passed into the chains. Simply writes error
// to writer. You should probably define your own to make something useful.
func defaultErrHandler(f EndpointFunc) http.HandlerFunc {
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
