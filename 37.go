package main

import "fmt"

func main() {

	var mp map[string]int = map[string]int{
		"apple":  5,
		"pear":   6,
		"orange": 9,
	}
	{
	}

	val, ok := mp["apple"]
	fmt.Println(val, ok)
	//fmt.Println(mp["apple"])
	//delete(mp, "apple")
	//mp["apple"] = 900

	//fmt.Println(mp)

}

/* MAPS allow us to store q value pairs*/
