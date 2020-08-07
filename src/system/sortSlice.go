package main

import (
	"fmt"
	"sort"
)

type aStructure struct {
	person string
	height int
	weight int
}

func main() {
	mySlice := []aStructure{}
	a := new(aStructure)
	a.person = "a"
	a.height = 1
	a.weight = 1
	b := new(aStructure)
	b.person = "b"
	b.height = 2
	b.weight = 2
	mySlice = append(mySlice, *a)
	mySlice = append(mySlice, *b)

	sort.Slice(mySlice, func(i, j int) bool {
		return mySlice[i].height < mySlice[j].height
	})
	fmt.Println("<:", mySlice)

	sort.Slice(mySlice, func(i, j int) bool {
		return mySlice[i].height > mySlice[j].height
	})
	fmt.Println(">:", mySlice)
}
