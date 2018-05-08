package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.ListenAndServe(":4000", Greeting("Release the gophers"))
}

type Greeting string

// ServerHTTP implements http.Handler
func (g Greeting) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, g)
}
