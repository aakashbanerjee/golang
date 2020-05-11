/*Using goroutines, create an incrementer program
have a variable to hold the incrementer value
launch a bunch of goroutines
each goroutine should
read the incrementer value
store it in a new variable
yield the processor with runtime.Gosched()
increment the new variable
write the value in the new variable back to the incrementer variable
use waitgroups to wait for all of your goroutines to finish
the above will create a race condition.
Prove that it is a race condition by using the -race flag*/

/*Fix the race condition you created in the previous exercise by using a mutex
it makes sense to remove runtime.Gosched()*/

package main

import (
	"fmt"
	"sync"
)

//Code below has a race condition
//go run -race lev9hoe3.go

func main() {
	counter := 0
	//Wg is a variable of type WaitGroup from the sync package
	var wg sync.WaitGroup
	//We will use mu which is a variable of type Mutex from the sync package to fix the race condition
	var mu sync.Mutex
	wg.Add(25)
	for i := 0; i < 25; i++ {
		go func() {
			//add mu.Lock() to avoid race
			mu.Lock()
			fmt.Println("Enter goroutine, counter is: ", counter)
			copyofcounter := counter
			//runtime.Gosched() not required when we are using mutex
			//runtime.Gosched()
			copyofcounter = copyofcounter + 1
			counter = copyofcounter
			fmt.Println("End goroutine, counter is: ", counter)
			//add mu.Unlock() to avoid race
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("All goroutines have ended: counter is: ", counter)
}
