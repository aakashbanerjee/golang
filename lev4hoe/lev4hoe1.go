/*Using a COMPOSITE LITERAL:
create an ARRAY which holds 5 VALUES of TYPE int
assign VALUES to each index position.
Range over the array and print the values out.
Using format printing
print out the TYPE of the array*/

package main

import "fmt"

func main() {
	myarr := [5]int{2, 4, 6, 8, 10}
	for i, v := range myarr {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", myarr)

	/*Using a COMPOSITE LITERAL:
	  create a SLICE of TYPE int
	  assign 10 VALUES
	  Range over the slice and print the values out.
	  Using format printing
	  print out the TYPE of the slice*/

	myslice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	for i, v := range myslice {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", myslice)

	/*Using the code from the previous example, use SLICING to create the following new slices which are then printed:
	  [42 43 44 45 46]
	  [47 48 49 50 51]
	  [44 45 46 47 48]
	  [43 44 45 46 47]*/

	mynewslice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	xi := mynewslice[:5]
	xj := mynewslice[5:]
	xk := mynewslice[2:7]
	xl := mynewslice[1:6]
	fmt.Println(xi)
	fmt.Println(xj)
	fmt.Println(xk)
	fmt.Println(xl)

	/*start with this slice
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	append to that slice this value
	52
	print out the slice
	in ONE STATEMENT append to that slice these values
	53
	54
	55
	print out the slice
	append to the slice this slice
	y := []int{56, 57, 58, 59, 60}
	print out the slice*/

	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52)
	fmt.Println(x)
	x = append(x, 53, 54, 55)
	fmt.Println(x)
	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	fmt.Println(x)

	/*start with this slice
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	use APPEND & SLICING to get these values here which you should ASSIGN to a variable “n” and then print:
	[42, 43, 44, 48, 49, 50, 51]*/

	m := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	n := []int{}
	n = append(n, m[:3]...)
	n = append(n, m[6:]...)
	fmt.Println(n)

	/*Create a slice to store the names of all of the states in the United States of America. What is the length of your slice? What is the capacity? Print out all of the values, along with their index position in the slice, without using the range clause. Here is a list of the states:
	` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`,
	*/

	states := make([]string, 50, 50)
	states = []string{` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`}
	fmt.Println("Length of States: ", len(states))
	fmt.Println("Capacity of States: ", cap(states))
	for i := 0; i < len(states); i++ {
		fmt.Printf("%d\t %s\n", i, states[i])
	}

	/*Create a slice of a slice of string ([][]string). Store the following data in the multi-dimensional slice:
	"James", "Bond", "Shaken, not stirred"
	"Miss", "Moneypenny", "Helloooooo, James."
	Range over the records, then range over the data in each record.*/

	xs1 := []string{"James", "Bond", "Shaken, not stirred"}
	xs2 := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	xs := [][]string{xs1, xs2}
	fmt.Println(xs)

	for i, v := range xs {
		fmt.Println(i, v)
		for j, k := range v {
			fmt.Println(j, k)
		}
	}

	/*Create a map with a key of TYPE string which is a person’s “last_first” name, and a value of TYPE []string which stores their favorite things. Store three records in your map. Print out all of the values, along with their index position in the slice.
	`bond_james`, `Shaken, not stirred`, `Martinis`, `Women`
	`moneypenny_miss`, `James Bond`, `Literature`, `Computer Science`
	`no_dr`, `Being evil`, `Ice cream`, `Sunsets`*/

	person := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	for k, v := range person {
		fmt.Println(k)
		for i, likes := range v {
			fmt.Printf("\nIndex %d: Value: %s", i, likes)
		}
		fmt.Println()
	}

	/*Using the code from the previous example, add a record to your map. Now print the map out using the “range” loop*/
	person["aakash"] = []string{`inspire`, `innovate`, `impact`}

	for k, v := range person {
		fmt.Println(k, v)
	}

	/*Using the code from the previous example, delete a record from your map. Now print the map out using the “range” loop*/
	delete(person, "aakash")
	for k, v := range person {
		fmt.Println(k, v)
	}
}
