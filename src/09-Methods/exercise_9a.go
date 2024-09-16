package main

import "fmt"

// type User struct {
// 	ID        int
// 	FirstName string
// 	LastName  string
// 	Email     string
// }

type Group struct {
	role           string
	users          []User
	newestUser     User
	spaceAvailable bool
}

// 1. Refactor the `describeUser` code to be a method (this should be a repetition of what was just completed in the course).
func (u *User) describe() string {
	desc := fmt.Sprintf("ID: %d\nName: %s %s\nEmail %s", u.ID, u.Lastname, u.Firstname, u.Email)
	return desc
}

// 2. Modify the `describeGroup` function to be a method called `describe()` that receives a `Group` type.
func (g *Group) describeGroup() string {
	desc := fmt.Sprintf("The %s user group has %d users. Newest user:  %s, Accepting New Users:  %t", g.role, len(g.users), g.newestUser.Firstname, len(g.users) < 2)
	return desc
}

func main() {
	u1 := User{ID: 1, Firstname: "Ibrahim", Lastname: "Adedokun", Email: "ibrahim@tiskae.com"}
	u2 := User{ID: 2, Firstname: "Rahmat", Lastname: "Adejumo", Email: "rahmat@tiskae.com"}

	u1Desc := u1.describe()
	u2Desc := u2.describe()

	// 3. Modify the main function to reflect those changes
	g := Group{role: "admin", users: []User{u1, u2}, newestUser: u2, spaceAvailable: false}
	gDesc := g.describeGroup()

	fmt.Println(u1Desc)
	fmt.Println(u2Desc)
	fmt.Println(gDesc)
}
