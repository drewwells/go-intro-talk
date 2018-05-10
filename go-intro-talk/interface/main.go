package main

import "fmt"

type Greeting string

func (h Greeting) String() string { // HL
	return fmt.Sprintf("Greeting, %s", string(h)) // HL
} // HL

func main() {
	fmt.Println(Greeting("Gophers"))
}
