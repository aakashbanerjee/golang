package main

import (
	"encoding/json"
	"fmt"
	"log"
)

/*Starting : with this code, marshal the []user to JSON. There is a little bit of a curve ball in
  this hands-on exercise - remember to ask yourself what you need to do to EXPORT a value from a package.*/
//User  exported
type user struct {
	First string //capital letter == exported
	Age   int
}

func main() {
	u1 := user{
		First: "James",
		Age:   32,
	}

	u2 := user{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := user{
		First: "M",
		Age:   54,
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	// your code goes here
	u, err := json.Marshal(users)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(u))

	v := []user{}

	err = json.Unmarshal(u, &v)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(v)

	s := "Aakash"
	b := []byte(s)

	fmt.Println(s)
	fmt.Println(b)
	fmt.Println(string(b))
}
