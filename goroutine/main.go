package main

import (
	"fmt"
	"runtime"
	"time"
)

func receiver(n string, ch <-chan int) {
	for {
		i, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(n, i)
	}
	fmt.Println(n, "is done.")
}

type Neko struct {
	Id int
}

func (n *Neko) String() string {
	return fmt.Sprintf("id: %d", n.Id)
}

func main() {
	go fmt.Println("yeah!")
	fmt.Printf("num cpu: %d\n", runtime.NumCPU())
	fmt.Printf("num goroutine: %d\n", runtime.NumGoroutine())
	fmt.Printf("num cpu: %s\n", runtime.Version())

	ch := make(chan int)

	go receiver("1st goroutine", ch)
	go receiver("2nd goroutine", ch)
	go receiver("3rd goroutine", ch)

	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(ch)
	time.Sleep(3 * time.Second)

	neko := Neko{Id: 1}
	fmt.Println(&neko)
}
