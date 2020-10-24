// map ma'lumot tuzilmasi haqida
// map[KeyType]ValueType

package main

import "fmt"

func main() {

	statuses := make(map[string]int)

	// map'ga qiymatlar qo'shish:
	statuses["pending"] = 0
	statuses["inited"] = 1
	statuses["running"] = 2
	statuses["timedout"] = 3
	statuses["successful"] = 4
	statuses["failed"] = 5

	// map'dan qiymatni o'qish:
	var successfulStatus = statuses["successful"]
	fmt.Println(successfulStatus)

	// map'dan bitta elementni o'chirib tashlash:
	delete(statuses, "timedout")
}
