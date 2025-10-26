package main

import (
	"fmt"

	"github.com/IndWit/golan"
)

func main() {
	fmt.Println("Welcome to Golan - A Go Test Project")
	fmt.Println("=====================================")
	fmt.Println()

	// Demonstrate the math functions
	a, b := 10, 5

	fmt.Printf("Adding %d + %d = %d\n", a, b, golan.Add(a, b))
	fmt.Printf("Subtracting %d - %d = %d\n", a, b, golan.Subtract(a, b))
	fmt.Printf("Multiplying %d * %d = %d\n", a, b, golan.Multiply(a, b))
	fmt.Printf("Dividing %d / %d = %d\n", a, b, golan.Divide(a, b))
}
