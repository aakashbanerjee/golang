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
}
