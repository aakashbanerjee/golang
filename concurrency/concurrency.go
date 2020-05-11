package main

import (
	"fmt"
	"runtime"
	"sync"
)

//Wg exported sync.WaitGroup
var Wg sync.WaitGroup

func main() {
	fmt.Println("GOOS: ", runtime.GOOS)
	fmt.Println("GOARCH: ", runtime.GOARCH)
	fmt.Println("CPUs: ", runtime.NumCPU())
	fmt.Println("Goroutines: ", runtime.NumGoroutine())

	Wg.Add(2)
	go funcA()
	go funcB()

	fmt.Println("CPUs: ", runtime.NumCPU())
	fmt.Println("Goroutines: ", runtime.NumGoroutine())
	Wg.Wait()
}

func funcA() {
	fmt.Println("In funcA")
	for i := 0; i < 10; i++ {
		fmt.Println("i: ", i)
	}
	Wg.Done()
}

func funcB() {
	fmt.Println("In funcB")
	for j := 0; j < 10; j++ {
		fmt.Println("j: ", j)
	}
	Wg.Done()
}
