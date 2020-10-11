// asosiy turlar (basic types)

package main

import "fmt"

func types() {
	// integer
	var X uint8 = 225
	fmt.Println(X+1, X)
	var Y int16 = 32765
	fmt.Println(Y - 2)

	// float
	a := 20.45
	b := 34.89
	c := b - a
	fmt.Printf("Natija: %f", c)
	fmt.Printf("\n c ning turi : %T", c)

	// complex
	var d complex128 = complex(6, 2)
	var e complex64 = complex(9, 2)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Printf("d ning turi %T va e ning turi %T", d, e)

	// boolean
	str1 := "VirtualDars"
	str2 := "virtualDars"
	str3 := "virtualdars"
	result1 := str1 == str2
	result2 := str1 == str3
	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Printf("result1 ning turi %T va result2 ning turi %T", result1, result2)

	// string
	str := "VirtualDars"
	fmt.Printf("\nstr matnning uzunligi %d", len(str))
	fmt.Printf("\nMatn: %s", str)
	fmt.Printf("\nstr ning turi: %T", str)
}
