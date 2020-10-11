package main

import "fmt"

func operators6() {
	a := 4

	// xotiradagi manzilini ishlatish (&) va ko'rsatkich (*) operatorlari
	b := &a
	fmt.Println(*b)
	*b = 7
	fmt.Println(a)
}
