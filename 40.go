package main

import "fmt"

func main() {

	var x map[string]int = map[string]int{"Hello": 3}
	//var x []int = []int{3, 4, 5}
	y := x
	y["y"] = 100
	y["7"] = 7
	fmt.Println(x, y)

	/*var x int = 5
	y := x
	y = 7

	fmt.Println(x, y)*/

}

/* mutable data types means its changable
immutable data types its inchangabel */
