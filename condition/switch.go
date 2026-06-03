package main

import (
	"fmt"
	"math/rand"
)

func switchstatement() {

outer:
	for {
		switch rnd := rand.Intn(100); {
		case rnd%2 == 0:
			fmt.Printf("generated number is %d which is even", rnd)
			break outer
		}
	}

}
