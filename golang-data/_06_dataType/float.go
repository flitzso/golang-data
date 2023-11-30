package main

import (
	"fmt"
)

func main() {
	// float32 Keyword
	var x float32 = 123.78
	var y float32 = 3.4e+38

	// float64 Keyword
	var z float64 = 1.7e+308

	fmt.Printf("Type: %T, value: %v\n", x, x)
	fmt.Printf("Type: %T, value: %v\n", y, y)
	fmt.Printf("Type: %T, value: %v", z, z)
}
