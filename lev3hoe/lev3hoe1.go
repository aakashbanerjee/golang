package main

import "fmt"

func main() {
	for num := 1; num <= 10000; num++ {
		fmt.Println("Num: ", num)
	}

	/*Print every rune code point of the uppercase alphabet three times. Your output should look like this:
	  65
	  	U+0041 'A'
	  	U+0041 'A'
	    U+0041 'A'
	  66
	  	U+0042 'B'
	  	U+0042 'B'
	  	U+0042 'B'
	   … through the rest of the alphabet characters*/
	for i := 65; i <= 90; i++ {
		fmt.Println(i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}
	}

	/*Create a for loop using this syntax for condition { }
	  Have it print out the years you have been alive.*/

	for year := 1982; year <= 2020; year++ {
		fmt.Println(year)
	}

	/*Create a for loop using this syntax for condition { }*/
	birthyear := 1982
	for {
		if birthyear > 2020 {
			break
		}
		fmt.Println(birthyear)
		birthyear++
	}

	//Print out the remainder (modulus) which is found for each number between 10 and 100 when it is divided by 4.

	for div := 10; div <= 100; div++ {
		printnum := div % 4
		fmt.Println(printnum)
	}

}
