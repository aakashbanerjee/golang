//Create a struct “customErr” which implements the builtin.error interface.
//Create a func “foo” that has a value of type error as a parameter.
//Create a value of type “customErr” and pass it into “foo”.

package main

import "fmt"

type customErr struct {
	info string
}

//method attached to type customErr that implements Error() from package builtin
func (ce customErr) Error() string {
	return fmt.Sprint("Error occurred...", ce.info)
}

func foo(e error) {
	fmt.Println(e)
}

func main() {
	customerror := new(customErr)
	customerror.info = "Here is a new error."
	foo(customerror)
}
