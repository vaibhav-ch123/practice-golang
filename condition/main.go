package main

import "fmt"

func reversedefer(ch rune) {
	fmt.Printf("%c", ch)
}

func main() {

	// var age = 10

	// if age := 10; age < 5 {
	// 	fmt.Println("ticket is free")
	// } else if age >= 5 && age <= 22 {
	// 	fmt.Println("ticket is $10")
	// } else {
	// 	fmt.Println("ticket is $15")
	// }

	// switchstatement()

	name := "vaibhav chauhan"

	for _, ch := range name {
		defer reversedefer(ch)
	}

}
