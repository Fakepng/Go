package main

import (
	"fmt"
	"strconv"
)

func FibonacciRecursion(n int) int {
	if n <= 1 {
		return n
	}
	return FibonacciRecursion(n-1) + FibonacciRecursion(n-2)
}

func main() {
	fmt.Print("Enter the number of terms: ")
	var terms int
	fmt.Scanln(&terms)

	for i := 0; i < terms; i++ {
		fmt.Print(strconv.Itoa(FibonacciRecursion(i)) + " ")
	}
	fmt.Print("\n")
	fmt.Print("Press any key to exit")
	var exit string
	fmt.Scanln(&exit)
}
