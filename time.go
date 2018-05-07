package main

import (
	"fmt"
	"time"
)

func main() {
	today, _ := time.Parse("Jan 2 2006", "May 10 2017")
	fmt.Println(today)
}
