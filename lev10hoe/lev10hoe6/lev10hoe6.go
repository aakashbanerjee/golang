/*write a program that
puts 100 numbers to a channel
pull the numbers off the channel and print them*/

package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan int)
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("Writing to channel c: ", i)
			c <- i
		}
		close(c)
		wg.Done()
	}()

	go func(c chan int) {
		for v := range c {
			fmt.Println("Reading and printing from channel c: ", v)
		}
		wg.Done()
	}(c)
	wg.Wait()
}
