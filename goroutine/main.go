package main

import (
	"fmt"
	"time"
)

func square(sqnumber int, sqch chan int) {

	num := 0

	for sqnumber != 0 {
		time.Sleep(100 * time.Millisecond)
		println("sq")
		digit := sqnumber % 10
		num += digit * digit * digit
		sqnumber /= 10
	}
	sqch <- num
}

func cube(cbnumber int, cbch chan int) {

	num := 0

	for cbnumber != 0 {
		time.Sleep(100 * time.Millisecond)
		println("cb")
		digit := cbnumber % 10
		num += digit * digit * digit
		cbnumber /= 10
	}

	cbch <- num
}

func write(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("successfully wrote", i, "to ch")
	}
	close(ch)
}

func main() {

	// number := 12345678
	// sqrch := make(chan int)
	// cbch := make(chan int)

	// go square(number, sqrch)
	// go cube(number, cbch)

	// sq, cb := <-sqrch, <-cbch
	// println(sq + cb)

	ch := make(chan int, 2)
	var ch1 chan int
	// go write(ch)
	// time.Sleep(2 * time.Second)
	// for v := range ch {
	// 	fmt.Println("read value", v, "from ch")
	// 	time.Sleep(2 * time.Second)

	// }
	fmt.Println(ch)
	fmt.Println(ch1)

}
