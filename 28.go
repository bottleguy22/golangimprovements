package main

import (
	"fmt"
)

func main() {

	name := "tim"
	fmt.Println("Before if")
	if name != "tim" {
		fmt.Println("inside if")

	}
	fmt.Println("After if")
}

func age() {

	age := 17

	if age >= 18 {
		fmt.Println("you can ride!")

	}
	fmt.Println("you cannot ride!")
}
