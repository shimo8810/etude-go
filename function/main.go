package main

import "fmt"

func main() {
	fmt.Println(plus(4, 5))

	f := func(x, y int) int { return x + y }
	fmt.Println(f(4, 5))

	fmt.Printf("%#v\n", func(x, y int) int { return x + y })
	fmt.Printf("%#v\n", func(x, y int) int { return x + y }(4, 5))

	f2 := returnFunc()
	f2()
	returnFunc()()

	callFunc(func() {
		fmt.Println("call function")
	})

	f3 := later()
	fmt.Println(f3("Golang"))
	fmt.Println(f3("is"))
	fmt.Println(f3("awesome"))

}

func plus(x, y int) int {
	return x + y
}

func returnFunc() func() {
	return func() {
		fmt.Println("function")
	}
}

func callFunc(f func()) {
	f()
}

func later() func(string) string {
	var store string

	return func(next string) string {
		s := store
		store = next
		return s
	}
}
