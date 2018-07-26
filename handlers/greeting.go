package handlers

import (
	"fmt"
	"io"
	"net/http"
)

// Greeter represents a type that is able to provide a greeting. This is an example of returning
// concretions and accepting interfaces.
type Greeter interface {
	Greet(name string) string
}

// GreetingHandler is a HTTP handler struct that handles greetings via a Greeter.
type GreetingHandler struct {
	greeter Greeter
}

// NewGreetingHandler returns a new GreetingHandler instance.
func NewGreetingHandler(greeter Greeter) *GreetingHandler {
	return &GreetingHandler{
		greeter: greeter,
	}
}

// ServeHTTP implements the http.Handler interface.
func (h *GreetingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	queryVals := r.URL.Query()

	name := ""
	if nameVal, ok := queryVals["name"]; ok {
		name = nameVal[0]
	}

	_, err := io.WriteString(w, h.greeter.Greet(name))
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to write response: %v", err), http.StatusInternalServerError)
	}
}
