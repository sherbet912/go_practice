package main

import (
	. "fmt"
	"math/rand"
	// "runtime"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		// runtime.Gosched()
		Println(s)
		sleep_time := rand.Intn(5)
		time.Sleep(time.Duration(sleep_time))
	}
}

func fibonacci(n int, c chan int) {
    x, y := 1, 1
    for i := 0; i < n; i++ {
        c <- x
        x, y = y, x + y
    }
    close(c)
}

func main() {
	// go say("helloooooo")
	// say("world!!!!")

	c := make(chan int, 10)
    go fibonacci(cap(c), c)
    for i := range c {
        Println(i)
    }
}