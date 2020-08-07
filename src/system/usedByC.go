package main

import "C"

import (
	"fmt"
)

// PrintMessage is a func
func PrintMessage() {
	fmt.Println("A Go function!")
}

// Multiply is a test
func Multiply(a, b int) int {
	return a * b
}

func main() {
	var a, b int = 1, 2
	Multiply(a, b)
	PrintMessage()
}
