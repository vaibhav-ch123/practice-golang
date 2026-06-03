package simpleinterest

import "fmt"

func init() {
	fmt.Println("Simple interest package initialization")
}

func Calculate(p, r, t float64) float64 {
	interest := p * (r / 100) * t
	return interest
}
