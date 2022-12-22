package main

import "fmt"

func main() {
	m := make(map[int]string)
	m[1] = "us"
	m[72] = "jp"
	fmt.Println(m)

	m2 := map[int]int{1: 1}
	v, ok := m2[2]
	fmt.Println(v, ok)

	m3 := map[int][]int{
		1: {1},
		2: {1, 2},
		3: {1, 2, 3},
	}
	s, ok := m3[4]
	fmt.Println(s, ok)

	for k, v := range m3 {
		fmt.Println(k, v)
	}
}
