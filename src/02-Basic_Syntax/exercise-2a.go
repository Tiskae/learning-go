package main

import "fmt"

func main() {
	// 1. In your `main` function, declare a variable that has a value of a sentence.
	sentence := "Mercedes and Audi I have not, I only drive BMWs"

	// 2. Iterate over that sentence using `range`.
	for pos, letter := range sentence {
		// 3. If the index of that letter is an odd number, print that letter.
		if pos%2 == 1 {
			fmt.Println(string(letter))
		}
		// 4. What do you notice?
	}

	// ## Hints

	// _Hint_: You might need to convert a type to get the output you expect

	// _Hint_: The modulo operator in Go functions similarly to
	// JavaScript.
}
