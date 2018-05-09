package main

import (
	"fmt"
	"net/http"
)

type Greeting string

// ServerHTTP implements http.Handler
func (g Greeting) ServeHTTP(w http.ResponseWriter, _ *http.Request) { // HL
	fmt.Fprint(w, g)
}

func main() {
	// http.ListenAndServer(addr string, handler http.Handler)
	http.ListenAndServe(":4000", Greeting("Release the gophers")) // HL
}
