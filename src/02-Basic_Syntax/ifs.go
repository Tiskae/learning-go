package main

import "fmt"

func ifs() {
	// if - else if - else
	score := 59

	if cutoff_mark := 50; score < cutoff_mark {
		fmt.Println("Sorry, you're not eligible for admission")
	} else if score == cutoff_mark {
		fmt.Println("Congratulations, just scaled through!")
	} else {
		fmt.Println("Bravo!, you are eligible for admission")
	}

	// error handing
	// if err:= hasError(): err != nil {
	// 	fmt.Println(err.Error())
	// }

}
