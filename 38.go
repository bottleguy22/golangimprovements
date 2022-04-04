package main

import "fmt"

//func add(x, y int) {
func add(x, y int) (z1 int, z2 int) {
	defer fmt.Println("hello")
	z1 = x + y
	z2 = x - y
	return
}

func main() {
	ans1, ans2 := add(14, 7)
	fmt.Println(ans1, ans2)

}

/* functions its a block of reusable code, you can reuse as many times you like
defer*/
