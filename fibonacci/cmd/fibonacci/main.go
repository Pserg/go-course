package main

import (
	"fmt"
	"pkg/fibonacci"
)

func main() {
	var input int
	fmt.Print("Enter the Number: ")
	fmt.Scanf("%d", &input)
	if input < 0 {
		fmt.Println("Number must be a positive")
		return
	}
	fmt.Println("Result", fibonacci.Calc(input))
}
