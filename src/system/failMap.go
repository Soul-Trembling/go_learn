package main

import "fmt"

func main() {
	aMap := map[string]int{}
	aMap["test"] = 1
	fmt.Println(aMap)

	/*
		error code
		aMap := make[string]int{}
		aMap = nil
		fmt.Println(aMap)
		aMap["test"] = 1
		fmt.Println(aMap)
	*/
}
