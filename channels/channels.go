//Simple example of a channel

package main

import "fmt"

func main() {
	c := make(chan int) //initialize the channel. remember channels block

	go func() {
		c <- 1982
	}()

	fmt.Println(<-c)
}
