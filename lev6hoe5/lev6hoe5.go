/*create a type SQUARE
create a type CIRCLE
attach a method to each that calculates AREA and returns it
circle area= π r 2
square area = L * W
create a type SHAPE that defines an interface as anything that has the AREA method
create a func INFO which takes type shape and then prints the area
create a value of type square
create a value of type circle
use func info to print the area of square
use func info to print the area of circle*/

package main

import (
	"fmt"
	"math"
	"strconv"
)

type square struct {
	length float64
	width  float64
	area   float64
}

type circle struct {
	radius float64
	area   float64
}

func (s square) calcArea() (ar float64) {
	return s.length * s.width
}

func (c circle) calcArea() (ar float64) {
	return math.Pi * c.radius * c.radius
}

type shape interface {
	calcArea() float64
}

func info(s shape) {
	fmt.Println("In func shape: ", s.calcArea())
}

func main() {
	fmt.Println("Good exercise to learn structs, methods, interfaces")
	//create a new SQUARE
	s1 := new(square)
	s1.length = 6.50
	s1.width = 6.50
	s1.area = s1.calcArea()
	fmt.Println("I am Square s1 with length: ", s1.length, "and width: ", s1.width, "and area: ", s1.area)
	//create a new CIRCLE
	c1 := new(circle)
	c1.radius = 4
	c1.area = c1.calcArea()
	fmt.Println("I am Circle c1 with radius: ", c1.radius, "and area: ", c1.area)

	info(s1)
	info(c1)

	//Build and use an anonymous func - hoe6
	func() {
		fmt.Println("I am in an anonymous func..")
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
	}()

	//Assign a func to a variable, then call that func
	f := func() {
		fmt.Println("I am in an anonymous func assigned to a variable")
		for i := 0; i < 10; i++ {
			fmt.Println("A" + strconv.Itoa(i))
		}
	}
	fmt.Println("Calling func using var f that has been initialized with a func")
	f()
	fmt.Printf("var f is of type: %T\n", f)

	/*hoe 8 Create a func which returns a func
	  assign the returned func to a variable
	  call the returned func*/
	returnedFunc := callA()
	fmt.Printf("var returnedFunc is of type: %T\n", f)
	returnedInt := returnedFunc()
	fmt.Println("I am the returned int from the returned func: ", returnedInt)

	//hoe 9
	/*A “callback” is when we pass a func into a func as an argument. For this exercise,
	  pass a func into a func as an argument -- NOT DONE passing a func inside a func as an argument*/

	//hoe10
	/*Closure is when we have “enclosed” the scope of a variable in some code block.
	  For this hands-on exercise, create a func which “encloses” the scope of a variable:*/
	fincrement := inc()
	fmt.Println(fincrement())
	fmt.Println(fincrement())
	fmt.Println(fincrement())
	fmt.Println(fincrement())
	fmt.Println(fincrement())
}

func inc() func() int {
	n := 0
	return func() int {
		n++
		return n
	}
}

func callA() func() int {
	return func() int {
		fmt.Println("I am in the return func")
		return 007
	}
}
