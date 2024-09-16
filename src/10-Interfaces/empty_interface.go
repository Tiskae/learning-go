package main

import "fmt"

type Users struct {
}

type Admin struct {
}

type Parent struct {
}

func main() {
	// empty interface - think "any" in TypeScript
	people := map[string]interface{}{
		"user":   Users{},
		"admin":  Admin{},
		"parent": Parent{},
	}

	fmt.Println(people)
}
