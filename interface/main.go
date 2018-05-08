package main

import "fmt"

func main() {
	fmt.Println(Greeting("Gophers"))
}

type Greeting string

func (h Greeting) String() string {
	return fmt.Sprintf("Greeting, %s", string(h))
}
