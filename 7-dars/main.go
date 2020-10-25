// funktsiyalar haqida

package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	var a int = 10
	var b int = 20
	fmt.Println(a, b)
	//swap(a, b)
	swapByref(&a, &b)
	fmt.Println(a, b)

	// result := sum(1, 2)
	// fmt.Println(result)
	/* result, err := sqrt(-5)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	} */
}

func swap(x, y int) int {
	var o int
	o = x
	x = y
	y = o
	return o
}

func swapByref(x, y *int) int {
	var o int
	o = *x
	*x = *y
	*y = o
	return o
}

func sum(x int, y int) int {
	return x + y
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Manfiy son uchun aniq emas")
	}
	return math.Sqrt(x), nil
}
