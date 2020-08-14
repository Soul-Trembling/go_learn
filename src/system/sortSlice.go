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
	a := aStructure{person: "a", height: 1, weight: 1}
	b := aStructure{"b", 2, 2}
	mySlice = append(mySlice, a)
	mySlice = append(mySlice, b)

	sort.Slice(mySlice, func(i, j int) bool {
		return mySlice[i].height < mySlice[j].height
	})
	fmt.Println("<:", mySlice)

	sort.Slice(mySlice, func(i, j int) bool {
		return mySlice[i].height > mySlice[j].height
	})
	fmt.Println(">:", mySlice)
}
