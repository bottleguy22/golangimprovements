package main

import (
	"fmt"
)

func main() {

	age := 19
	/*
		if age >= 18 {
			fmt.Println("You can ride")
		} else if age >= 14 {
			fmt.Println("You can ride with a parent!")
		} else {
			fmt.Println("You cannot ride kid!!")
		}
	*/
	if age >= 18 {
		fmt.Println("you can ride alone!")
	}
	if age >= 14 && age < 18 {
		fmt.Println("you can ride with a parent!")
	}
	if age < 14 {
		fmt.Printf("you cannot ride!")
	}
}

/*
If
Else if
Else
*/
