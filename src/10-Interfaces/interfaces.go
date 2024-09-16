package main

import "fmt"

type User struct {
	ID        int
	Firstname string
	Lastname  string
	Email     string
}

type Group struct {
	role           string
	users          []User
	newestUser     User
	spaceAvailable bool
}

type Describer interface {
	describe() string
}

func (u *User) describe() string {
	desc := fmt.Sprintf("ID: %d\nName: %s %s\nEmail %s", u.ID, u.Lastname, u.Firstname, u.Email)
	return desc
}

func (g *Group) describe() string {
	desc := fmt.Sprintf("The %s user group has %d users. Newest user:  %s, Accepting New Users:  %t", g.role, len(g.users), g.newestUser.Firstname, len(g.users) < 2)
	return desc
}

// Do the describing
func doTheDescribing(d Describer) string {
	return d.describe()
}

func interfaces() {
	u1 := User{ID: 1, Firstname: "Ibrahim", Lastname: "Adedokun", Email: "ibrahim@tiskae.com"}
	u2 := User{ID: 2, Firstname: "Rahmat", Lastname: "Adejumo", Email: "rahmat@tiskae.com"}

	u1Desc := doTheDescribing(&u1)
	u2Desc := doTheDescribing(&u2)

	g := Group{role: "admin", users: []User{u1, u2}, newestUser: u2, spaceAvailable: false}
	gDesc := doTheDescribing(&g)

	fmt.Println(u1Desc)
	fmt.Println(u2Desc)
	fmt.Println(gDesc)
}
