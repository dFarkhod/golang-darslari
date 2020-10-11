package main

import "fmt"

func operators4() {
	a := 14
	b := 30

	// & (bitwise AND)
	result1 := a & b
	fmt.Printf("a & b ning natijasi %d", result1)

	// | (bitwise OR)
	result2 := a | b
	fmt.Printf("\np | q ning natijasi %d", result2)

	// ^ (bitwise XOR)
	result3 := a ^ b
	fmt.Printf("\np ^ q nint natijasi %d", result3)

	// << (left shift)
	result4 := a << 1
	fmt.Printf("\np << 1 ning natijasi %d", result4)

	// >> (right shift)
	result5 := a >> 1
	fmt.Printf("\np >> 1 ning natijasi %d", result5)

	// &^ (AND NOT)
	result6 := a &^ b
	fmt.Printf("\np &^ q ning natijasi %d", result6)
}
