package main

import "fmt"

type Greeting string

func (h Greeting) String() string {
	return fmt.Sprintf("Greeting, %s", string(h))
}

func main() {
	fmt.Println(Greeting("Gophers"))
}

var _ fmt.Stringer = Greeting("") // HL
