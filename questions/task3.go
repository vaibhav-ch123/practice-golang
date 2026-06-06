package main

import (
	"fmt"
	"sync"
)

func displaynumber(wg *sync.WaitGroup, num int) {

	defer wg.Done()

	fmt.Println(num)

}

func printnumber() {

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go displaynumber(&wg, i)
	}

	wg.Wait()
}
