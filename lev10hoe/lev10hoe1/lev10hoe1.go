/*get this code https://play.golang.org/p/j-EA6003P0 working:
using func literal, aka, anonymous self-executing func
solution: https://play.golang.org/p/SHr3lpX4so
a buffered channel
solution: https://play.golang.org/p/Y0Hx6IZc3U */

package main

import (
	"fmt"
)

func main() {
	c := make(chan int) //general non-buffered channel
	//c := make(chan int, 1) //general buffered channel of size 1
	//using anonymous self executing function

	//if using buffered channel no need to use self-execution anonymous func
	go func() {
		c <- 42
	}()
	//put the c<-42 in the self-executing func which is also a go routine.
	//c <- 42

	fmt.Println(<-c)
}
