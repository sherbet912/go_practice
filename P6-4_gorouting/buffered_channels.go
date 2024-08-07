package main

import (
	"fmt"
)

func main() {
	msg := make(chan string, 2)

	msg <- "A"
	msg <- "B"
	fmt.Println(<-msg)
	fmt.Println(<-msg)
}