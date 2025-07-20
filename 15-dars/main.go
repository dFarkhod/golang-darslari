package main

import "fmt"

func Filter[T any](items []T, predicate func(T) bool) []T {
	var result []T
	for _, item := range items {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}

type User struct {
	Name string
	Age  int
}

func main() {
	evens := Filter([]int{1, 2, 3, 4, 5}, isEven)
	fmt.Println(evens)

	adults := Filter([]User{
		{"Alice", 17},
		{"Bob", 21},
		{"Carol", 15},
	}, func(u User) bool {
		return u.Age >= 18
	})

	fmt.Println(adults)
}

func isEven(n int) bool {
	return n%2 == 0
}
