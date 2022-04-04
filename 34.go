package main

import "fmt"

func main() {

	var a []int = []int{5, 6, 7, 8, 9}
	b := append(a, 10)

	fmt.Println(b)

	/*fmt.Println(cap(a[:3]))
	fmt.Println(cap(a))

	var x [5]int = [5]int{1, 2, 3, 4, 5}

	var s []int = x[1:3]
	fmt.Println(s[:cap(s)])
	fmt.Println(s[1:])*/

}

/* slices its a portion of an array
slice = fatia
capacity = how many elements are left on the array
length   capacity*/
