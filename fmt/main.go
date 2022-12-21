package main

import "fmt"

func main() {
	fmt.Println("Hello, Golang!")
	fmt.Println("My", "name", "is", "Taro")

	fmt.Printf("10進法=%d 2進法=%b 8進法=%o 16進法=%x\n", 17, 17, 17, 17)
	fmt.Printf("%d年%d月%d日\n", 2022, 12)
	fmt.Printf("%d年%d月%d日\n", 2022, 12, 15, 72)

	fmt.Printf("数値=%v 文字列=%v 配列=%v\n", 5, "Golang", [...]int{1, 2, 3})
	fmt.Printf("数値=%#v 文字列=%#v 配列=%#v\n", 5, "Golang", [...]int{1, 2, 3})
	fmt.Printf("数値=%T 文字列=%T 配列=%T\n", 5, "Golang", [...]int{1, 2, 3})

	// 標準エラーに出力される
	print("hello, world!")
	println("hello, world!")

}
