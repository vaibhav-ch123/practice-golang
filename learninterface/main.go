package main

import "fmt"

type new_interface1 interface {
	function1()
	function2()
}

type new_interface2 interface{}

type new_struct1 struct {
	field1 int
	field2 string
}

func (s new_struct1) function1() {
	fmt.Println("implemented function1 by value receiver")
}

func (s *new_struct1) function2() {
	fmt.Println(("implemented funtion2 by pointer receiver"))
}

func main() {

	// var s1 new_struct1
	// var i1 new_interface1 = &s1
	// fmt.Printf("type of interface %T and value %v\n", i1, i1)
	// fmt.Printf("type of structure %T and value %v\n", s1, s1)
	// s1.function1()
	// s1.function2()

	s2 := new_struct1{4, "hello"}
	s2.function1()
	s2.function2()

	// var i2 interface{}
	// fmt.Printf("type of interface %T and value %v\n", i2, i2)
}
