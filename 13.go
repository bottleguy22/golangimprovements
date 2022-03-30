package main

import (
	"fmt"
)

func main() {
	var width int
	var height int

	fmt.Println("Whats the measure of the width?")
	fmt.Scanf("%d", &width)

	fmt.Println("Whats the measure of the height?")
	fmt.Scanf("%d", &height)

	perimeter := 2 * (width + height)
	fmt.Println(perimeter)
}
