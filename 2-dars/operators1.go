// arifmetik operatorlar

package main

import "fmt"

func operators1() {
	a := 64
	b := 20

	// qo'shuv
	result1 := a + b
	fmt.Printf("a + b ning natijasi = %d", result1)

	// ayruv
	result2 := a - b
	fmt.Printf("\na - b ning natijasi = %d", result2)

	// ko'paytiruv
	result3 := a * b
	fmt.Printf("\na * b ning natijasi = %d", result3)

	// bo'luv
	result4 := a / b
	fmt.Printf("\na / b nint natijasi = %d", result4)

	// modul
	result5 := a % b
	fmt.Printf("\na %% b ning natijasi = %d", result5)
}
