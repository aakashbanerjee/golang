/*Create a pup package. The pup package should have an exported func “Years”
which takes dog's age in years and turns them into dog's age in human years (1 dog year = 7 human years).
Document  your code with comments. Use this code in func main.
run your program and make sure it works
run a local server with godoc and look at your documentation.*/

package main

import (
	"fmt"

	"github.com/aakashbanerjee/golang/lev12hoe/pup"
)

func main() {
	dogAge := 7
	fmt.Println("Dog's Age: ", dogAge)
	humanYears := pup.Years(dogAge)
	fmt.Println("If age of dog is: ", dogAge, " age of dog in human years is: ", humanYears)
}
