package main

import "fmt"

func main() {
	aSliceLiteral := []int{1, 2, 3, 4, 5}
	integer := make([]int, 20)
	for i := 0; i < len(integer); i++ {
		fmt.Println(i)
	}

	aSliceLiteral = nil
	integer = append(integer, -5900)
	s2 := integer[1:3]
	fmt.Println(s2, aSliceLiteral)

	s1 := make([]int, 5)
	reSlice := s1[1:3]
	fmt.Println(s1)
	fmt.Println(reSlice)
	reSlice[0] = -100
	reSlice[1] = 123456
	fmt.Println(s1)
	fmt.Println(reSlice)
}
