/*Create your own type “person” which will have an underlying type of “struct” so that it can store the following data:
first name
last name
favorite ice cream flavors
Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice which stores
the favorite flavors. */

package main

import "fmt"

func main() {

	type person struct {
		firstname string
		lastname  string
		flavors   []string
	}

	p1 := person{
		firstname: "Aakash",
		lastname:  "Banerjee",
		flavors: []string{
			"chocolate",
			"hazelnut",
			"vanilla"},
	}

	p2 := person{
		firstname: "Ryan",
		lastname:  "Banerjee",
		flavors: []string{
			"strawberry",
			"bubblegum",
			"peanut butter",
		},
	}

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(p1.firstname, p1.lastname, p1.flavors)
	fmt.Println()
	fmt.Printf("%s's favorite icecreams flavors are: \n", p1.firstname)
	for _, v := range p1.flavors {
		fmt.Printf("%s \t\n", v)
	}

	fmt.Println(p2.firstname, p2.lastname, p2.flavors)
	fmt.Println()
	fmt.Printf("%s's favorite icecreams flavors are: \n", p2.firstname)
	for _, v := range p2.flavors {
		fmt.Printf("%s \t\n", v)
	}

	/*Take the code from the previous exercise, then store the values of type person in a map with the key of last name.
	  Access each value in the map. Print out the values, ranging over the slice.*/

	m1 := map[string]person{
		p1.firstname: p1,
		p2.firstname: p2,
	}

	fmt.Println(m1)
	fmt.Println("Printing values in map m1:")
	for _, v1 := range m1 {
		fmt.Println(v1.firstname, v1.lastname)
		fmt.Println()
		for _, flavor := range v1.flavors {
			fmt.Printf("%s\t", flavor)
		}
		fmt.Println("\n--------")
	}

	/*Create a new type: vehicle.
	  The underlying type is a struct.
	  The fields:
	  doors
	  color
	  Create two new types: truck & sedan.
	  The underlying type of each of these new types is a struct.
	  Embed the “vehicle” type in both truck & sedan.
	  Give truck the field “fourWheel” which will be set to bool.
	  Give sedan the field “luxury” which will be set to bool. solution
	  Using the vehicle, truck, and sedan structs:
	  using a composite literal, create a value of type truck and assign values to the fields;
	  using a composite literal, create a value of type sedan and assign values to the fields.
	  Print out each of these values.
	  Print out a single field from each of these values.*/

	type vehicle struct {
		doors int
		color string
	}

	type truck struct {
		vehicle
		fourWheel bool
	}

	type sedan struct {
		vehicle
		luxury bool
	}

	t1 := truck{
		vehicle: vehicle{
			doors: 4,
			color: "black",
		},
		fourWheel: true,
	}

	s1 := sedan{
		vehicle: vehicle{
			doors: 2,
			color: "yellow",
		},
		luxury: true,
	}

	fmt.Println(t1)
	fmt.Println(s1)
	fmt.Println("Color of Truck: ", t1.color)
	fmt.Println("Car doors: ", s1.doors)

	//Create and use an anonymous struct.

	s := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "Aakash",
		friends: map[string]int{
			"Cesar": 111111111,
			"Roya":  222222222,
			"Ryan":  333333333,
		},
		favDrinks: []string{"coors", "bud"},
	}

	fmt.Println(s)
	fmt.Println(s.first)
	fmt.Println(s.friends)
	fmt.Println(s.favDrinks)

	for k, _ := range s.friends {
		fmt.Println(k)
	}
}
