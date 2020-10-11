package main

import "fmt"

func operators5() {
	var a int = 35
	var b int = 60

	a = b
	fmt.Println(a)

	a += b
	fmt.Println(a)

	a -= b
	fmt.Println(a)

	a *= b
	fmt.Println(a)

	a /= b
	fmt.Println(a)

	a %= b
	fmt.Println(a)
}
