package main

import "fmt"

// Create a function called `average` that takes three arguments separate arguments, all of which are floats.
// The function should return the average of the three arguments as a float.
func average(num1, num2, num3 float64) (avg float64) {
	total := num1 + num2 + num3
	avg = total / 3.0
	return
}

// Refactor your code to use a variadic function that takes in an unknown number of arguments.
func average_variadic(floats ...float64) (avg float64) {
	var total float64

	for _, value := range floats {
		total += value
	}

	avg = total / float64(len(floats))
	return
}

func main() {
	avg := average(4.33, 4.75, 4.57)
	avg2 := average_variadic(4.1, 4.2, 4.3, 4.4, 4.5, 4.6, 4.7, 4.8, 4.9, 5.0)
	fmt.Println(avg, avg2)

}
