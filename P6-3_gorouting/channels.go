package main

import (
	"fmt"
	"time"
)

func add(c chan int, num int, t int) {
	time.Sleep(100 * time.Millisecond)
	// 接收channel資料
	a := <-c
	a = a + num
	fmt.Println(t, "-num's value:", a)
	// 發送channel資料
	c <- a
}

func main() {
	c := make(chan int)
	num := 10

	go add(c, 10, 1)
	c <- num  //發送channel資料
	num = <-c // 接收channel資料
	fmt.Println(num)

	go add(c, 10, 2)
	c <- num  //發送channel資料
	num = <-c // 接收channel資料
	fmt.Println(num)

	go add(c, 10, 3)
	c <- num  //發送channel資料
	num = <-c // 接收channel資料
	fmt.Println(num)

	close(c)
}