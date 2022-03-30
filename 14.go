package main

import (
	"fmt"
)

func main() {
	var planet string
	var isTrue bool
	var temp float64

	planet, isTrue, temp = "Mars", true, 19.5
	fmt.Println("Air is good on", planet)
	fmt.Println("Its", isTrue)
	fmt.Println("It is", temp, "degrees")
}
