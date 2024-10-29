package main

import (
	"time"
)

func worker(id int, jobs <-chan int, result chan<- int) {
	for i := range jobs {
		println("worker", id, "started job: ", i)
		time.Sleep(time.Second)
		println("worker", id, "finished job: ", i)
		result <- i * 2
	}
}

func main() {
	//宣告工作數目
	const numJobs = 5
    
    //STEP 2:
	// 建立 jobs channel 
	jobs := make(chan int, numJobs)
	// 建立 results channel 
	results := make(chan int, numJobs)
    
    //STEP 3:
	// creates 3 goroutines for the worker function
	for w := 1; w <= 3; w++ {

		go worker(w, jobs, results)
	}
    
    //STEP 4:
	// 發送工作至 jobs channel
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
    
    //STEP 5:
	// 關閉 jobs channel 
	close(jobs)

    //STEP 6:
	// 從 results channel接收結果
	for a := 1; a <= numJobs; a++ {
		<-results
	}
    
    //STEP 8:
	// 關閉  results channel 
	close(results)
}