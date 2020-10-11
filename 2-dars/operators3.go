// mantiqiy operatorlar

package main

import "fmt"

func operators3() {
	var a int = 14
	var b int = 30

	if a != b && a <= b {
		fmt.Println("true")
	}

	if a != b || a <= b {
		fmt.Println("true")
	}

	if !(a == b) {
		fmt.Println("true")
	}

}
