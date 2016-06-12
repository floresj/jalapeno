# jalapeno
A tiny set of commonly used funcs/patterns that I constantly find myself copying from project to project.

## Purpose
Avoid me copying the same crap over and over. Now I can just import it!

## The Gist
Ability to easily define middleware and endpoints that return an error. These errors are then managed by a single `ErrorHandler` so that endpoints are concerned with handling controller logic and delegate error handling to a specific error handler.

## Examples
```go
// Create an object where we define an ErrorHandler that will be used to handle all errors captured be all 
// Endpoint() calls as well as any number of middleware funcs
c := web.NewChain(ErrorHandler, Logger)

m := http.NewServeMux()

// For your individual handlers, pass in an Endpoint. An endpoint returns a HandlerFunc
m.HandleFunc("/users", c.Endpoint(UserHandler))

// I return an error if something goes wrong and let someone else deal with it
func UserHandler(w http.ResponseWriter, r *http.Request) error {
	// do stuff
	u, err := SaveUser(r)
	
	// Oh crap, something went wrong
	if err != nil {
	// This will be handled by the ErrorHandler defined below
	return err
	}
	
	// Do something with u if no error
}

// I handle stuff when things go wrong. If you have custom errors defined,
// you can handle them here
func ErrorHandler(f web.EndpointFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := f(w, r)
		if err == nil {
			return
		}
		switch err {
		case SomePreDefinedErr:
		    http.Error(w, err.Error(), http.StatusInternalServerError)
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
```

## TODO
- Docs docs docs
- Add more stuff
