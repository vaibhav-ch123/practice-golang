package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var jobs = make(chan Jobs, 10)
var results = make(chan Results, 10)

type Jobs struct {
	job_id        int
	random_number int
}

type Results struct {
	worker_id     int
	job_id        int
	random_number int
	sumofdigit    int
}

func allocate_job(noofjob int) {

	for i := 1; i <= noofjob; i++ {
		random_number := rand.Intn(999)
		job := Jobs{i, random_number}
		jobs <- job
	}

	close(jobs)
}

func worker_pool(noofworker int) {

	var wg sync.WaitGroup
	for i := 1; i <= noofworker; i++ {
		wg.Add(1)
		go workers(i, &wg)
	}

	wg.Wait()

	close(results)
}

func workers(worker_id int, wg *sync.WaitGroup) {

	for job := range jobs {
		result := Results{worker_id, job.job_id, job.random_number, digitSum(job.random_number)}
		results <- result

	}
	wg.Done()
}

func outputResult(op chan bool) {

	for result := range results {
		fmt.Println("worker id: ", result.worker_id, ", ", "job id: ", result.job_id, ", ", "random number: ", result.random_number, ", ", "sumofdigit: ", result.sumofdigit)
	}

	op <- true
}

func digitSum(number int) int {

	sumofdigit := 0

	for number != 0 {
		sumofdigit += number % 10
		number /= 10
	}

	time.Sleep(2 * time.Second)
	return sumofdigit
}

func main() {

	starttime := time.Now()

	go worker_pool(50)

	go allocate_job(100)

	var op chan bool = make(chan bool)
	go outputResult(op)

	<-op

	endtime := time.Now()
	diff := endtime.Sub(starttime)

	println("total time taken to complete jobs: ", diff.Seconds(), " second")

}
