package main

import "fmt"

func sum(num1, num2 int) (result int) {
	result = num1 + num2
	return
}

// Variadic function
func average(numbers ...int) (avg int) {
	total := 0
	for _, value := range numbers {
		total += value
	}
	avg = total / len(numbers)
	return
}

func main() {
	money_left := sum(255, 200)
	fmt.Println(money_left)

	avg := average(1, 2, 3, 4, 5)
	fmt.Println(avg)
}
