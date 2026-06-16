package main

import "fmt"

func main() {
	var x interface{} = 42
	y, ok := x.(string)
	fmt.Println(y, ok)
}
