package main

import "fmt"

type User struct {
	ID                         int
	Firstname, Lastname, Email string
}

// Method
func (u *User) Describe() string {
	desc := fmt.Sprintf("ID: %d\nName: %s %s\nEmail %s", u.ID, u.Lastname, u.Firstname, u.Email)
	return desc
}

// Functional
func describe(u *User) string {
	desc := fmt.Sprintf("ID: %d\nName: %s %s\nEmail %s", u.ID, u.Lastname, u.Firstname, u.Email)
	return desc
}

func methods() {
	user1 := User{ID: 005, Firstname: "Ibrahim", Lastname: "Adedokun", Email: "ibrahim@bendingspoons.com"}

	descFunctional := describe(&user1)
	descMethod := user1.Describe()

	fmt.Println(descFunctional)
	fmt.Println(descMethod)
}
