package main

import (
	"fmt"
	"time"
)

func evennumber(wait chan bool) {

	for i := 2; i <= 10; i += 2 {
		fmt.Println(i)
		time.Sleep(2000 * time.Millisecond)
	}

	wait <- true
}

func oddnumber(wait chan bool) {

	for i := 1; i <= 10; i += 2 {
		fmt.Println(i)
		time.Sleep(2000 * time.Millisecond)
	}

	wait <- true
}

func firsthalfsum(arr []int, result chan int) {

	sum := 0
	for _, val := range arr {
		sum += val
	}
	result <- sum
}

func secondhalfsum(arr []int, result chan int) {

	sum := 0
	for _, val := range arr {
		sum += val
	}
	result <- sum
}

func download(str string) {

	fmt.Printf("Downloading %s...\n", str)
	time.Sleep(3 * time.Second)

	fmt.Printf("Completed %s\n", str)
}

func printevenodd() {

	// wait := make(chan bool)

	// go evennumber(wait)
	// go oddnumber(wait)

	// <-wait

	// result := make(chan int)

	// arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	// go firsthalfsum(arr[0:4], result)
	// go secondhalfsum(arr[4:8], result)

	// firsthalf := <-result
	// secondhalf := <-result

	// fmt.Println("sum of arr ", firsthalf+secondhalf)

	go download("file1")
	go download("file2")
	go download("file3")

	time.Sleep(4 * time.Second)
}
