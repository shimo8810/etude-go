package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	if x, y := 1, 2; x < y {
		fmt.Printf("x(%d) is less than y(%d)\n", x, y)
	}

	arr := [...]int{1, 2, 4, 7, 11}

	for i, n := range arr {
		fmt.Println(i, n)
	}

	str := "いろはにほへと"

	for i, r := range str {
		fmt.Printf("[%02d] %c\n", i, r)
	}

	switch x := 12; x {
	case 1, 2:
		fmt.Println("1 or 2")
	case 3, 4:
		fmt.Println("3 or 4")
	default:
		fmt.Println("unknown")
	}

	var x interface{} = 3.14
	switch x.(type) {
	case bool:
		fmt.Println("bool")
	case float32, float64:
		fmt.Println("float")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("don't know")
	}

}
