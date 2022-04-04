package main

import "fmt"

func returnFunc(x string) func() {

	return func() { fmt.Println(x) }

}

/*func test2(myFunc func(int) int) {

fmt.Println(myFunc(7))
}*/

func main() {

	returnFunc("hello")()
	x := returnFunc("goodbye")
	x()
	/*func() {
		fmt.Println("test")
	}()

	test := func(x int) int {
		return x * -1
		fmt.Println("Hello!")
	}
	test3 := func(x int) int {
		return x * 7
	}
	test2(test3)
	test(8)
	fmt.Println(test)

	test(5)*/

}

//function advanced and function closure
