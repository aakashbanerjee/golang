package main

import "fmt"

type varaakash int

var myvar varaakash

func main() {
	fmt.Println(myvar)
	fmt.Printf("%T\n", myvar)
	myvar = 42
	fmt.Println(myvar)
}
