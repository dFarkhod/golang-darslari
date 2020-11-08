// Goroutine'lar haqida

package main

import (
	"fmt"
	"time"
)

func main() {
	go display("salom berdik")
	go display("alik oldik")

	fmt.Scanln()
}

func display(input string) {
	for i := 1; true; i++ {
		fmt.Println(i, input)
		time.Sleep(1 * time.Second)
	}
}
