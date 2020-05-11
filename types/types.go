//make keyword is for slices, maps and channels only | returns initialized value of type T not a pointer *T
//new keyword works on any type T to allocate memory and zero it. new does not initialize memory | returns *T

package main

import (
	"fmt"
	"os"
	"strconv"
)

type friend struct {
	name string
	loc  string
}

func main() {
	fmt.Println("Hello types!")
	slice1 := new([]int)
	slice2 := new([]friend)
	fmt.Println("Slice1: ", slice1)         //slice is &[]
	fmt.Printf("Slice1 Type: %T\n", slice1) //slice type is a pointer *[]int
	fmt.Println("Slice1 Pointer: Before", *slice1)
	//adding data to the slice
	for i := 0; i < 10; i++ {
		*slice1 = append(*slice1, i)
	}

	fmt.Println("Slice1 Pointer: After", *slice1)

	f := new(friend)
	f.name = "David"
	f.loc = "New York"

	fmt.Println("Struct friend: ", *f)

	*slice2 = append(*slice2, *f)

	fmt.Println("Slice of Type Struct Slice2: ", *slice2)

	for j := 0; j < 10; j++ {
		newFriend := new(friend)
		newFriend.name = "A" + strconv.Itoa(j)
		newFriend.loc = "B" + strconv.Itoa(j)
		*slice2 = append(*slice2, *newFriend)
		fmt.Println("Slice2: ", *slice2)
	}

	for i, v := range *slice2 {
		fmt.Println(i, v)
		fmt.Println(i, v.name)
		fmt.Println(i, v.loc)
	}

	fmt.Println("Capacity Slice2: ", cap(*slice2))
	fmt.Println("Length Slice2: ", len(*slice2))

	f1 := new(friend)
	f1.name = "Ryan"
	f1.loc = "Lathrop"
	*slice2 = append(*slice2, *f1)

	fmt.Println("Capacity Slice2: ", cap(*slice2))
	fmt.Println("Length Slice2: ", len(*slice2))
	fmt.Println("Slice2: ", *slice2)

	//create a new slice of type friend and copy slice2 into the new slice.
	slice3 := new([]friend)
	*slice3 = append(*slice3, *slice2...)
	fmt.Println("Slice3: ", *slice3)

	//create a new map with key of type string and value of type (slice of struct friend)
	m1 := make(map[string][]friend)
	fmt.Println(m1)

	for i := 0; i < 10; i++ {
		key := "ken" + strconv.Itoa(i)
		fmt.Println(key)
		m1[key] = *slice2
		fmt.Println(m1)
	}

	//example forever loop and breaking out of it
	getout := false
	for {
		if getout == true {
			break
		}
		fmt.Println(m1["ken2"])
		getout = true
	}

	//os
	home := os.Getenv("HOME")
	fmt.Println(home)
	goroot := os.Getenv("GOROOT")
	fmt.Println(goroot)
	gopath := os.Getenv("GOPATH")
	fmt.Println(gopath)
	user := os.Getenv("USER")
	fmt.Println(user)
}
