package main

import (
	"fmt"
	"sync"
)

var count int = 0

func inc(wg *sync.WaitGroup, m *sync.Mutex) {

	defer wg.Done()

	m.Lock()
	count = count + 1
	m.Unlock()
}

func count_increment() {

	var wg sync.WaitGroup
	var m sync.Mutex
	for i := 1; i <= 10000; i++ {
		wg.Add(1)
		go inc(&wg, &m)
	}

	wg.Wait()

	fmt.Println(count)

}
