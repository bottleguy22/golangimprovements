package main

import (
	"fmt"
)

func main() {

	//age := 17
	age := 19

	if age >= 18 {
		fmt.Println("you can ride!")

	} else {
		fmt.Println("you cannot ride!")
		fmt.Printf("Wait %d years\n", 18-age)
	}
}

/*
If
Else if
Else
*/
