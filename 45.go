package main

import "fmt"

type Point struct {
	x int32
	y int32
}

func changeX(pt *Point) {

	pt.x = 100
	fmt.Println(pt)
}

func main() {

	//p1 := Point{x: 3}
	p1 := &Point{y: 3}
	fmt.Println(p1)
	changeX(p1)
	fmt.Println(p1)

	//fmt.Println(p1)

	//var p1 Point = Point{1, 2, false}
	//var p2 Point = Point{-5, 7, true}
	//p1.x = 7

	//fmt.Println(p1.y, p2.y)
	//fmt.Println(p1.y)

}

//struts and custom types
