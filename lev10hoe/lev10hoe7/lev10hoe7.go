/*write a program that
launches 10 goroutines
each goroutine adds 10 numbers to a channel
pull the numbers off the channel and print them*/

package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	grs := 10
	//var wg sync.WaitGroup

	//wg.Add(11)

	for i := 0; i < grs; i++ {
		go func() {
			for j := 0; j < 10; j++ {
				fmt.Println("Writing to channel c: ", j)
				c <- j
			}
			//	wg.Done()
		}()
	}

	for k := 0; k < 100; k++ {
		fmt.Println("Reading and printing from channel c: ", k, <-c)
	}

	//wg.Wait()

}
