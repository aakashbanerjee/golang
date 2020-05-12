/*Create a dog package. The dog package should have an exported func “Years”
which takes human years and turns them into dog years (1 human year = 7 dog years).
Document  your code with comments. Use this code in func main.
run your program and make sure it works
run a local server with godoc and look at your documentation.*/

package lev12hoe

import (
	"fmt"

	"github.com/aakashbanerjee/golang/lev12hoe/dog"
)

func main() {
	humanAge := 38
	dogAge := dog.Years(humanAge)
	fmt.Println("If age of dog in human years is: ", humanAge, " age of dog in dog years is: ", dogAge)
}
