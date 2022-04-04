package main

import "fmt"

func main() {

	arr := [3]int{4, 5, 6}
	sum := 0

	arr2D := [2][3]int{{1, 2}, {3, 4}}
	//an array inside another array that have 2 elements and 2 arrays
	fmt.Println(arr2D[0][2])
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	fmt.Println(sum)
	/*fmt.Println(len(arr))

	var arr [5]int
	[0 0 0 0 0]
	 0 1 2 3 4
	arr[0] = 100
	[100 0 0 0 0]
	arr[4] = 900
	[100 0 0 0 900]
	var arr [5]string
	fmt.Println(arr[0]) */

}

/*arrays = collection of elements
[5] - 5 items its the only value that he can handle, if a change will be more
the index is always zero
len = LENGTH = COMPRIMENTO
its possible to have another arrays inside arrays {{3,4}, {5,6},{7,8}}*/
