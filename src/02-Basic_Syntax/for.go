package main

import "fmt"

func for_loop() {
	fmt.Println("Bomb! Time is ticking")

	// for timer := 100; timer >= 0; timer-- {
	// 	fmt.Println(timer)
	// }

	timer := 100

	for timer >= 0 {
		fmt.Println(timer)
		timer--
	}

	fmt.Println("Boom!")

	for pos, char := range "I LOVE YOU" {
		fmt.Println("Position: ", pos, "Character: ", string(char))
	}

}
