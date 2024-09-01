package main

import "fmt"

func maps() {
	// var userEmails = make(map[int]string)

	// userEmails[1] = "rahmat@tiskae.io"
	// userEmails[2] = "ibrahim@tiskae.io"

	userEmails := map[int]string{
		1: "rahmat@tiskae.io",
		2: "ibrahim@tiskae.io",
	}

	userEmails[1] = "morenikeji@tiskae.io"

	if firstEmail, ok := userEmails[1]; ok {
		fmt.Printf("User %d exists with the email: %s\n", 1, firstEmail)
	} else {
		fmt.Printf("User %d does not exist\n", 1)
	}

	userEmails[3] = "imposter@tiskae.io"
	delete(userEmails, 3)

	fmt.Println(userEmails, len(userEmails))
}
