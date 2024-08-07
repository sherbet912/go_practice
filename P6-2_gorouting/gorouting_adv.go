package main

import (
	"fmt"
	"sync"
)

func main() {
	var wait_group sync.WaitGroup
	// 2 個 gorouting
	wait_group.Add(2)

	go func() {
		count("oranges")
		wait_group.Done()
	}()

	go func() {
		count("apples")
		wait_group.Done()
	}()

	count("banana")
	wait_group.Wait()

}

func count(thing string) {
	for i := 0; i < 4; i++ {
		fmt.Printf("counting %s\n", thing)
	}
}