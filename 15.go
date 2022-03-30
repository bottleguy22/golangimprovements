package main

import "fmt"

func multi() (int, int) {
	return 5, 4
}

func main() {

	b := 4
	_, b = multi()
	fmt.Println(b)

}
