package main

import "fmt"

func sumoftwo() {

	var a, b int
	fmt.Scanln(&a, &b)

	fmt.Println(a + b)
}

func swaptwonum() {

	a, b := 5, 4
	a, b = b, a
	fmt.Println(a, b)
}

func wordCount() {

	var num int
	fmt.Print("Choose number of words: ")
	fmt.Scanln(&num)

	var word string
	mp := make(map[string]int)

	for i := 0; i < num; i++ {
		fmt.Scanln(&word)
		mp[word]++
	}

	for wrd, val := range mp {
		fmt.Printf("%s freq: %d\n", wrd, val)
	}
}
