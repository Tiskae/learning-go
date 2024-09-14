package main

import "fmt"

// ## Setup

// - Reference the code discussed in `structs.md` and `code/structs.go`
type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
}

type Group struct {
	role           string
	users          []User
	newestUser     User
	spaceAvailable bool
}

// - Let's say that in `describeGroup`, we only want to accept new users (represented by the `spaceAvailable` field) if there are fewer than 2 existing users in the group.
func describeGroup(g Group) string {
	desc := fmt.Sprintf("The %s user group has %d users. Newest user:  %s, Accepting New Users:  %t", g.role, len(g.users), g.newestUser.FirstName, len(g.users) < 2)
	return desc
}

func main() {
	u1 := User{ID: 1, FirstName: "Ibrahim", LastName: "Adedokun", Email: "ibrahim@tiskae.com"}
	u2 := User{ID: 2, FirstName: "Rahmat", LastName: "Adejumo", Email: "rahmat@tiskae.com"}

	g := Group{role: "admin", users: []User{u1, u2}, newestUser: u2, spaceAvailable: false}
	gDesc := describeGroup(g)

	fmt.Println(gDesc)
}

// ## Directions

// 1. In the `describeGroup`, update `spaceAvailable` to be `false` if there are 2 or more users.

// 2. Reprint the variable `g` at the end of `main()`. What do you notice?
