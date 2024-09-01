package main

import "fmt"

func arraysAndSlices() {
	// var scores [5]float64 = [5]float64{1, 2, 3, 4, 5}
	// var scores = [5]float64{1, 2, 3, 4, 5}
	// var scores = [5]float64{1, 2, 3, 4, 5}
	scores := [5]float64{1, 2, 3, 4, 5}
	fmt.Println(scores)
	// scores2 = make()

	// Slices
	fruitsArray := [7]string{"banana", "apple", "orange", "pineapple", "pawpaw", "kiwi", "mango"}
	slicedFruits := fruitsArray[2:6]
	fruitsToAdd := append(slicedFruits, "pear", "cherry")

	fmt.Println(fruitsArray, len(fruitsArray), cap(fruitsArray))
	fmt.Println(slicedFruits, len(slicedFruits), cap(slicedFruits))
	fmt.Println(fruitsToAdd, len(fruitsToAdd), cap(fruitsToAdd))
}
