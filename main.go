package main

import "fmt"

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

func main() {
	fmt.Println(totalAndResult(1, 2, 3, 4, 5))
}
