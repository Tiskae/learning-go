package main

import "fmt"

func doThings() {
	defer fmt.Println("First line?")
	defer fmt.Println("Second line?")
	fmt.Println("Will this come last?")
}

func recoverFromPanic() {
	if r := recover(); r != nil {
		fmt.Println("We panicked but everything will be alright")
		fmt.Println(r)
	}
}

func doOtherThings() {
	defer recoverFromPanic()

	for i := 0; i < 5; i++ {
		fmt.Println(i)

		if i == 2 {
			panic("PANIC!")
		}
	}
}

func main() {
	// Defer
	doThings()

	// Recover
	doOtherThings()

	fmt.Println("Function execution continues...")
}
