package main

import "fmt"

func main() {

	var arr [5]int

	//[0 0 0 0 0]
	// 0 1 2 3 4
	arr[0] = 100
	arr[4] = 900
	fmt.Println(arr[0])
}
