package main

import (
	"fmt"
	"learnstructs/computer"
)

func main() {
	spec := computer.Spec{
		Maker: "apple",
		Price: 50000,
	}

	fmt.Println(spec)
}
