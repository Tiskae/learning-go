package main

import (
	"fmt"
)

// 2. Write a function that calculates and returns the average score (also a float)
//   - Use the `range` keyword
func averageScore(scores [7]float64) (average float64) {
	var total float64

	for _, value := range scores {
		total += value
	}

	average = total / float64(len(scores))
	return
}

// 2. Write a function that takes a string argument and returns a boolean indicating whether or not that key exists in your map of pets.
func mapHasKey(mp map[string]string, key string) (exists bool) {
	_, exists = mp[key]
	return
}

// 2. Write a function that takes one or more groceries as strings and appends them to the slice, printing out the resulting list of groceries.
func addItemsToSlice(oldSlice []string, items ...string) {
	newSlice := append(oldSlice, items...)
	fmt.Println(newSlice)
}

func main() {
	// 	## Directions

	// ### Part 1

	// 1. Instantiate an array of scores
	//     - The array should have at least 5 elements of type `float64`
	scores := [7]float64{61.7, 69.8, 75.3, 85.2, 96.7, 51.1, 70.0}
	average := averageScore((scores))

	fmt.Printf("The average of %v is %f\n", scores, average)

	// ### Part 2

	// 1. Define a map that contains a set of pet names, and their corresponding animal type. i.e.: `"fido": "dog"`.
	petNames := map[string]string{
		"tom":       "cat",
		"spike":     "dog",
		"wabbit":    "rabbit",
		"whisperer": "parrot",
	}

	fmt.Println(mapHasKey(petNames, "tom"), mapHasKey(petNames, "jerry"))

	// ### Part 3

	// 1. Instantiate a slice that has an initial value of a collection of groceries.
	groceries := []string{"garri", "epa", "sugar", "kulikuli"}
	addItemsToSlice(groceries, "milk", "turkey")
}
