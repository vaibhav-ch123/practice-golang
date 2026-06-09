package main

import (
	"fmt"
)

func removeelement(arr []int, position int) []int {

	return append(arr[0:position], arr[position+1:]...)

}

func insertelement(arr []int, position int, val int) []int {

	newarr := append(arr[0:position], val)
	return append(newarr, arr[position:]...)
}

func duplicateelement(arr []int) {

	freq := make(map[int]int)

	for _, val := range arr {
		freq[val]++
	}

	for val, count := range freq {
		if count > 1 {
			fmt.Println(val)
		}
	}
}

func leftshift(arr []int, shift int) []int {

	return append(arr[shift:], arr[0:shift]...)
}

func rightshift(arr []int, shift int) []int {

	return append(arr[len(arr)-shift:], arr[0:len(arr)-shift]...)
}
