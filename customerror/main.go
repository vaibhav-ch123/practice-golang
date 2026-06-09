package main

import (
	"customerror/errorhandling"
	"fmt"
)

func circleArea(radius float64) (float64, error) {
	if radius < 0 {
		return 0, errorhandling.New(radius)
	}
	return radius * 3.14, nil
}

func main() {

	area, err := circleArea(-3)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(area)
}
