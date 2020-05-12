//We are going to learn about testing next.
//For this exercise, take a moment and see how much you can figure out about testing by reading the
//testing documentation & also Caleb Doxsey’s article on testing.
//See if you can get a basic example of testing working.

package lev11hoe5

import "fmt"

func main() {
	a := 6
	b := 4
	total := sum(a, b)
	fmt.Printf("Sum of %v, %v is %v\n", a, b, total)
}

func sum(a int, b int) int {
	return a + b
}
