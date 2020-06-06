package main

import "fmt"

type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	foods := []string{"pizza", "potato", "water"}
	me := person{age: 34, name: "test", favFood: foods}
	fmt.Println(me)
}

func totalAndResult(numbers ...int) (int, bool) {
	total := 0
	result := false

	for _, number := range numbers {
		total += number
	}

	if total > 0 {
		result = true
	}

	return total, result
}
