package main

import "fmt"

func changeValue(str *string) {
	*str = "changed!"
}

func changeValue2(str string) string {
	str = "changed!"
	return str
}
func main() {

	toChange := "hello"
	fmt.Println(toChange)
	changeValue2(toChange)
	//changeValue(&toChange)
	fmt.Println(toChange)

	//x := 7
	//y := &x

	//fmt.Println(x, y)
	//*y = 8
	//fmt.Println(x, y)

	//fmt.Println(&x)

	//fmt.Println(x)

}

//pointers e dereference operator (& and *)
