package main

import (
	"fmt"
	"sync"
)

func recoverfrompanic() {

	if r := recover(); r != nil {
		fmt.Println("recovered from: ", r)
	}
}

func panicfunction(val int) {

	defer recoverfrompanic()

	panic(fmt.Sprintf("%d is divisible by 7 which is a panic", val))
}

func worker(wg *sync.WaitGroup, job chan int, result chan int) {

	defer wg.Done()

	for val := range job {
		if val%7 == 0 {
			panicfunction(val)
		} else {
			result <- val * val
		}
	}

}

func assignjob(noofjob int, job chan int) {

	for i := 1; i <= noofjob; i++ {
		job <- i
	}

	close(job)
}

func displayresult(result chan int) {

	for val := range result {
		fmt.Println(val)
	}
}

func goroutinepanicandrecovery() {

	job := make(chan int)
	result := make(chan int, 20)
	var wg sync.WaitGroup

	go assignjob(20, job)

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(&wg, job, result)
	}

	wg.Wait()

	close(result)

	displayresult(result)

}
