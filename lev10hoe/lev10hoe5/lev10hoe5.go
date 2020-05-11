//Show the comma ok idiom starting with this code - https://play.golang.org/p/YHOMV9NYKK.

package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 1) //buffered channel

	c <- 42

	//comma, ok idiom
	v, ok := <-c
	fmt.Println(v, ok)

	close(c)

	v, ok = <-c
	fmt.Println(v, ok)
}
