// struct haqida

package main

import "fmt"

type car struct {
	Make, Model, Color string
	Year, Weight       int
	Engine             engine
}

type engine struct {
	name string
	hp   int
}

func main() {
	//var myCar Car
	myCar := car{Make: "Volvo", Model: "XC90", Color: "White", Year: 2020, Weight: 2343, Engine: engine{"T8", 400}}
	fmt.Println(myCar)
}
