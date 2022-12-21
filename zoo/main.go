package main

import (
	"fmt"

	"zoo/animals"
)

func main() {
	fmt.Println(AppName())

	fmt.Println(animals.FeedElephant())
	fmt.Println(animals.FeedMonkey())
	fmt.Println(animals.FeedRabbit())

}
