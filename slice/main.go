package main

import "fmt"

func main() {

	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(arr)

	// fmt.Println(removeelement(arr, 5))
	// fmt.Println(insertelement(arr, 5, 5))

	// arr1 := []int{0, 1, 2, 0, 2, 4, 3, 2, 1}
	// duplicateelement(arr1)

	fmt.Println(leftshift(arr, 3))
	fmt.Println(rightshift(arr, 3))

}
