package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func main() {

	var dataset = `[
{"color": "red", "quantity": 1},
{"color": "blue", "quantity": 2}
]`

	type Shoe struct {
		Color    string
		Quantity int
	}

	dec := json.NewDecoder(strings.NewReader(dataset))
	var shoes []Shoe
	if err := dec.Decode(&shoes); err != nil {
		fmt.Println(err)
	}

	for _, shoe := range shoes {
		fmt.Printf("%+v\n", shoe)
	}
}
