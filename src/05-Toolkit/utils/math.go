package utils

import "fmt"

func printNum(num int) {
	fmt.Println("Current number: ", num)
}

// Adds multiple numbers together
func Add(numbers ...int) (total int) {
	total = 0

	for _, v := range numbers {
		printNum(total)
		total += v
	}

	return
}
