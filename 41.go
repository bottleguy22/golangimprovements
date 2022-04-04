package main

import "fmt"

func main() {

	var x [2]int = [2]int{3, 4}
	y := x
	y[0] = 100

	//var x map[string]int = map[string]int{"Hello": 3}
	//var x []int = []int{3, 4, 5}
	//y := x
	//y["y"] = 100
	//x["7"] = 7
	fmt.Println(x, y)

	//var x int = 5
	//y := x
	//y = 7
	//fmt.Println(x, y)

}

//mutable e immutable data types
