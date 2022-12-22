package main

import "fmt"

func main() {

	// slice
	arr1 := make([]int, 5, 10)
	arr1[0] = 1
	fmt.Printf("%v %#v %T\n", arr1, arr1, arr1)
	fmt.Println(len(arr1), cap(arr1))

	// slice
	arr2 := []int{1, 2, 3, 4, 5}
	fmt.Printf("%v %#v %T\n", arr2, arr2, arr2)

	// array
	arr3 := [...]int{1, 2, 3, 4, 5}
	fmt.Printf("%v %#v %T\n", arr3, arr3, arr3)

}
