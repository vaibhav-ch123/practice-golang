package main

import (
	"fmt"
	_ "fmt"
	"polymorphism/rectangle"
	"polymorphism/square"
	"reflect"
)

type shapes interface {
	Area() float64
	Perimeter() float64
}

func display(slc []shapes) {

	for _, shape := range slc {
		fmt.Printf("Area of %v: %.2f and perimeter: %.2f\n", reflect.TypeOf(shape).Name(), shape.Area(), shape.Perimeter())
	}
}

func main() {

	rec1 := rectangle.New(5.4, 6.2)
	sqr1 := square.New(4.5)

	slc := []shapes{rec1, sqr1}

	display(slc)

}
