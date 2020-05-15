//An example how interfaces work in Go.

//create abstractions by considering the functionality that is common between datatypes,
//instead of the fields that are common between datatypes
//an interface{} value is not of any type; it is of interface{} type
//interfaces are two words wide; schematically they look like (type, value)
//it is better to accept an interface{} value than it is to return an interface{} value
//a pointer type may call the methods of its associated value type, but not vice versa
//everything is pass by value, even the receiver of a method
//an interface value isn’t strictly a pointer or not a pointer, it’s just an interface
//if you need to completely overwrite a value inside of a method, use the * operator to manually dereference a pointer

package main

import "fmt"

type cat struct {
}

type dog struct {
}

type horse struct {
}

//Remember: Any type that has a method speak() is also of type animal
//like cat, dog, horse
type animal interface {
	speak() string
}

func (c *cat) speak() string {
	return "meow!"
}

func (d *dog) speak() string {
	return "woof!"
}

func (h *horse) speak() string {
	return "neigh!"
}

func main() {
	fmt.Println("I am a demo of Go interfaces.")
	iamacat := new(cat)
	iamadog := new(dog)
	iamahorse := new(horse)
	//animals is a variable that is a slice of type interface animal.
	//the slice holds a var of type cat, dog and horse each
	//since eat of the type cat, dog and horse has a method speak attached to it, they satisfy the interface speak
	//depending on what type we call speak with the corresponding method for that type will be invoked.
	animals := []animal{iamacat, iamadog, iamahorse}
	for _, a := range animals {
		fmt.Println(a.speak())
	}

	a1 := new(animal)
	//type: *main.animal > a1 is a pointer to the type animal which is of interface type in package main
	fmt.Printf("Type of var a1 initialized as a new variable of type interface animal: %T\n", a1)
	//some experiments
	*a1 = iamacat
	fmt.Printf("Type of var a1 assigned iamacat which is *main.cat: %T\n", a1)
	fmt.Printf("var a1 assigned iamacat which is *main.cat: %T", *a1)
}
