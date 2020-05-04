package main

import "fmt"

const (
	myconstantTyped   bool = true
	myconstantUntyped      = 4.77
	rawstring              = `Hello
  I am a raw string
  I do not like the raw string literal`
)

const (
	a = 2017 + iota
	b
	c
	d
)

func main() {
	num := 17
	fmt.Printf("%d\t%b\t%#x\n", num, num, num)
	fmt.Printf("myconstantTyped: %T %v\t myconstantUntyped: %T %v\t\n", myconstantTyped, myconstantTyped, myconstantUntyped, myconstantUntyped)
	fmt.Println(rawstring)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
