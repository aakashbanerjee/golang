/*This exercise will reinforce our understanding of method sets:
create a type person struct
attach a method speak to type person using a pointer receiver
*person
create a type human interface
to implicitly implement the interface, a human must have the speak method
create func “saySomething”
have it take in a human as a parameter
have it call the speak method
show the following in your code
you CAN pass a value of type *person into saySomething
you CANNOT pass a value of type person into saySomething
here is a hint if you need some help*/

package main

import "fmt"

//Person  is a structure to store name and age of a person
type Person struct {
	name string
	age  int
}

func (p *Person) speak() {
	fmt.Println("Hello my name is: ", p.name)
	fmt.Println("Hello I am: ", p.age, " years old")
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	iamaperson := new(Person) //iamaperson is of type *Person
	iamaperson.age = 7
	iamaperson.name = "Ryan"
	saySomething(iamaperson)
}
