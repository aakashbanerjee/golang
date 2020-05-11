package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func main() {
	/*Create a value and assign it to a variable.
	  Print the address of that value.*/

	val := 1982
	fmt.Println("Address where val is stored: ", &val)

	//create a new variable of type person
	p := new(person)
	fmt.Printf("Type of p %T\n", p)
	fmt.Println("Address of p:", &p)
	(*p).name = "Aakash"
	(*p).age = 38
	fmt.Println("p (Before changeMe is called):", *p)
	p.changeMe()
	fmt.Println("p (After changeMe is called):", *p)
}

func (p *person) changeMe() {
	p.name = "Roya"
	p.age = 42
}
