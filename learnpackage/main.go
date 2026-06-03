package main

import (
	"fmt"
	"learnpackage/simpleinterest"
	"log"
)

var p, r, t = 5000.0, 10.0, 1.0

func init() {
	fmt.Println("Main package initialized")
	if p < 0 {
		log.Fatal("Principal is less than zero")
	}
	if r < 0 {
		log.Fatal("Rate of interest is less than zero")
	}
	if t < 0 {
		log.Fatal("Duration is less than zero")
	}
}

func main() {

	si := simpleinterest.Calculate(p, r, t)
	fmt.Println("Simple interest is", si)

}
