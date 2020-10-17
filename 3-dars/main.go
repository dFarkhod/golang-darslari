package main

import (
	"fmt"
	"time"
)

func main() {
	testSwitch3()
}

func testIf() {
	i := 7
	if i == 7 {
		fmt.Println("Yetti!")
	}
}

func testIfElse() {
	points := 20000
	if points < 5000 {
		fmt.Println("Silver")
	} else if points < 10000 {
		fmt.Println("Gold")
	} else {
		fmt.Println("Platinum")
	}
}

func testSwitch1() {
	weekday := time.Now().Weekday()
	switch weekday {
	case 1:
		fmt.Println("Dushanba")
	case 2:
		fmt.Println("Seshanba")
	case 3:
		fmt.Println("Chorshanba")
	case 4:
		fmt.Println("Payshanba")
	case 5:
		fmt.Println("Juma")
	case 6:
		fmt.Println("Shanba")
	case 7:
		fmt.Println("Yakshanba")
	default:
		fmt.Println("Xato")
	}
}

func testSwitch2() {
	greeting := 1
	switch {
	case greeting == 1:
		fmt.Println("Assalomu alaykum")
	case greeting == 2:
		fmt.Println("أسلم عليكم")
	case greeting == 3:
		fmt.Println("Hello")
	case greeting == 4:
		fmt.Println("Bonjour")
	case greeting == 5:
		fmt.Println("Привет")
	default:
		fmt.Println("Xato")
	}
}

func testSwitch3() {
	var userChoice string = "uch"
	switch userChoice {
	case "bir":
		fmt.Println("C#")
	case "ikki", "uch":
		fmt.Println("Go")
	case "to'rt", "besh", "olti":
		fmt.Println("TypeScript")
	}
}
