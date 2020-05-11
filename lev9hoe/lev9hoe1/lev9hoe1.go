/*in addition to the main goroutine, launch two additional goroutines
each additional goroutine should print something out
use waitgroups to make sure each goroutine finishes before your program exists*/

package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

//Wg variable WaitGroup
var Wg sync.WaitGroup

func main() {
	start := time.Now()
	Wg.Add(2)
	go funcA()
	go funcB()
	fmt.Println("Number of goroutines executing: ", runtime.NumGoroutine())
	Wg.Wait()
	fmt.Println("Time elapsed goroutine main: ", time.Since(start))
}

func funcA() {
	start := time.Now()
	for i := 0; i < 20; i++ {
		fmt.Println("Value of i from funcA: ", i)
	}
	fmt.Println("Time elapsed goroutine funcA: ", time.Since(start))
	Wg.Done()
}

func funcB() {
	start := time.Now()
	for i := 0; i < 20; i++ {
		fmt.Printf("Value of i from funcB: %v Type of i from fyncB: %T\n", i, i)
	}
	fmt.Println("Time elapsed goroutine funcB: ", time.Since(start))
	Wg.Done()
}
