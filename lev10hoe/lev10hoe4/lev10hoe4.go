/*Starting with this code https://play.golang.org/p/MvL6uamrJP, pull the values off the channel using a select statement*/

package main

import (
	"fmt"
)

func main() {
	q := make(chan int)
	c := gen(q)

	receive(c, q)

	fmt.Println("about to exit")
}

func gen(q chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		q <- 1
		close(c)
	}()
	return c
}

func receive(c <-chan int, q <-chan int) {
	for {
		select {
		case v := <-c:
			fmt.Println("Value from channel c: ", v)
		case <-q:
			return
		}
	}
}
