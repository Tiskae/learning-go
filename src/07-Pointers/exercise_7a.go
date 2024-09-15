package main

import "fmt"

type User struct {
	ID                         int
	FirstName, LastName, Email string
}

// 2. Write a function called `updateEmail` that takes in a `*User` type
func updateEmail(u *User, newEmail string) {
	// 3. Update the user's email to something new
	u.Email = newEmail
}

func main() {

	// ## Directions

	// 1. Define an instance of the User struct
	user1 := User{ID: 1, FirstName: "Cristiano", LastName: "Ronaldo", Email: "cristianoronaldo@gmail.com"}

	// 4. Call `updateEmail()` from `main()` and verify the updated email has persisted
	updateEmail(&user1, "cristianoronaldo@alnassr.com")

	fmt.Println(user1)

	// _HINT_ Check out the [go docs](https://tour.golang.org/moretypes/4) if things feel a little weird
}
