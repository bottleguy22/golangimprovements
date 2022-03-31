package main

import "fmt"

func main() {

	/*x := 0

	for x <= 5 {
		fmt.Println(x)
		x++

	}*/
	/*for x := 0; x <= 5; x += 2 {
		//break
		continue
		//fmt.Println(x)
	}*/
	for x := 0; x <= 1000; x++ {
		if x != 0 && x%3 == 0 && x%7 == 0 && x%9 == 0 {
			fmt.Println(x)
			continue
		} /*else{
		fmt.Println("N")
		}*/
	}
}

//start using some loops
//incrementing 1 you use ++
//-- subtractive one
