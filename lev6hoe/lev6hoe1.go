package main

import (
	"fmt"
)

/*Hands on Exercise 4
Create a user defined struct with
the identifier “person”
the fields:
first
last
age
attach a method to type person with
the identifier “speak”
the method should have the person say their name and age
create a value of type person
call the method from the value of type person*/

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("In method speak():")
	fmt.Println("My firstname is: ", p.first)
	fmt.Println("My lastname is: ", p.last)
	fmt.Println("My age is: ", p.age)
}

func main() {
	//Use the “defer” keyword to show that a deferred func runs after the func containing it exits.
	defer printSalud()
	valueFoo := foo()
	valueBarint, valueBarstring := bar()
	fmt.Println("ValueFoo: ", valueFoo)
	fmt.Println("ValueBar: ", valueBarint, valueBarstring)
	s1 := new([]int)
	for i := 1; i <= 10; i++ {
		*s1 = append(*s1, i)
	}
	sums1 := foo1(*s1...) //s1 unfurled, passing to foo1 that accepts variadic parameter
	fmt.Println("Sum of s1: ", sums1)
	s2 := make([]int, 30)
	for j := 0; j < 20; j++ {
		s2[j] = j
	}
	fmt.Println("s2: ", s2)
	sums2 := bar1(s2)
	fmt.Println("Sum of s2: ", sums2)

	//create a new person
	newPerson := new(person)
	newPerson.first = "Aakash"
	newPerson.last = "Banerjee"
	newPerson.age = 38
	newPerson.speak()
}

/*Hands on exercise #1
create a func with the identifier foo that returns an int
create a func with the identifier bar that returns an int and a string
call both funcs
print out their results*/

//func (r receiver) identifier (parameters) (return(s)){code}

func foo() int {
	return 10
}

func bar() (int, string) {
	return 10, "Aakash"
}

/*Hands on exercise #2
create a func with the identifier foo1 that
takes in a variadic parameter of type int
pass in a value of type []int into your func (unfurl the []int)
returns the sum of all values of type int passed in
create a func with the identifier bar1 that
takes in a parameter of type []int
returns the sum of all values of type int passed in*/

func foo1(xi ...int) int {
	fmt.Println("xi: ", xi)
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}

func bar1(xi []int) int {
	fmt.Println("xi: ", xi)
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}

func printSalud() {
	fmt.Println("I am a Ninja - Level 6 in Golang")
}
