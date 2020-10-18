package main

import (
	"fmt"
	"sort"
)

func main() {
	testSlices3()
}

func testArrays1() {
	// myarr deb nomlangan va 3ta string turidagi
	// elementlardan iborat bo'lgan yangi qator yaratamiz:
	var myarr [3]string

	// elementga qiymat berishda uning indeksi yoziladi
	// va indekslash 0dan boshlanadi:
	myarr[0] = "GO"
	myarr[1] = "dasturlash"
	myarr[2] = "VirtualDars"

	// elementning qiymatini olishda ham indeks ishlatiladi:
	fmt.Println("Qatorning elementlari:")
	fmt.Println("1: ", myarr[0])
	fmt.Println("2: ", myarr[1])
	fmt.Println("3: ", myarr[2])
}

func testArrays2() {
	// qatorni e'lon qilishning qisqa uslubi:
	myarr := [3]int{2, 4, 8}
	fmt.Println(myarr)
}

func testArrays3() {
	// qatorlarni solishtirish
	myarr1 := [...]int{9, 7, 6, 5}
	myarr2 := [4]int{9, 7, 6, 5}
	fmt.Println(myarr1 == myarr2)
}

func testArrays4() {
	// ko'p o'lchamli qatorlar:
	myarr := [3][3]string{{"C#", "GO", "TypeScript"},
		{"Java", "C++", "Python"},
		{"Dart", "Kotlin", "Swift"}}
	fmt.Println(myarr[0][1])
}

func testArrays5() {
	myarr1 := [3]int{2, 4, 8}

	// qatordan to'liq nusxa olish:
	myarr2 := myarr1
	fmt.Println(myarr1)
	fmt.Println(myarr2)

	myarr1[2] = 100
	fmt.Println(myarr1)
	fmt.Println(myarr2)
}

func testArrays6() {
	myarr1 := [3]int{2, 4, 8}

	// qatordan reference'li nusxa olish:
	myarr2 := &myarr1
	fmt.Println(myarr1)
	fmt.Println(*myarr2)

	myarr1[2] = 100
	fmt.Println(myarr1)
	fmt.Println(*myarr2)
}

func testSlices1() {
	// yangi slice e'lon qilish:
	myslice := []int{2, 4, 8}

	// slice oxiriga yangi element qo'shish:
	myslice = append(myslice, 10)

	// slice'ni uzunligini olish
	fmt.Printf("slice'ning uzunligi: %d", len(myslice))
}

func testSlices2() {
	myarr := [7]string{"olma", "anor", "shaftoli", "gilos",
		"o'rik", "uzum", "anjir"}

	// slice'ni qatordan yaratib olish:
	myslice := myarr[1:4]
	fmt.Println(myslice)
}

func testSlices3() {
	// slice'ni tartiblash:
	myslice := []int{45, 67, 23, 90, 33, 21, 56, 78, 89}
	sort.Ints(myslice)
	fmt.Println(myslice)
}
