// for loop haqida

package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	// 1-uslub, oddiy for loop'i:
	for i := 1; i < 8; i++ {
		fmt.Println(i)
	}

	// 2-uslub, infinite (checksiz) loop:
	for {
		fmt.Println("Alhamdulillah " + strconv.Itoa(time.Now().Second()))
		time.Sleep(1 * time.Second)
		if time.Now().Second() == 59 {
			break // break - for loop'ni to'htatib, dasturni ishlashi undan chiqadi
		}

	}

	// 3-uslub, while loop'iga o'shab:
	j := 1
	for j < 10 {
		fmt.Println(j)
		j++
	}

	// 4-uslub, for range yordamida to'plamlar bo'ylab yurib chiqish

	// qator elementlari bo'ylab for range orqali yurib chiqish:
	myarr := [3]string{"yer", "quyosh", "oy"}
	for index, value := range myarr {
		fmt.Println("index:", index, "value:", value)
	}

	// map elementlari bo'ylab for range orqali yurib chiqish:
	mymap := map[int]string{
		1: "c#",
		2: "typescript",
		3: "go",
	}
	for key, value := range mymap {
		fmt.Println("key:", key, "value:", value)
	}
}
