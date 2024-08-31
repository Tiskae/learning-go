package main

import "fmt"

func switches() {
	// Switch
	grade := "A"
	var remark string

	switch grade {
	case "A":
		remark = "Excellent"
	case "B":
		remark = "Good"
	case "C":
		remark = "Fair"
	case "D":
		remark = "Poor"
	case "E":
		remark = "Very poor"
	default:
		remark = "Fail"
	}

	fmt.Println(grade, remark)

	number := 59
	switch {
	case number > 49:
		fmt.Println("You passed")
		fallthrough
	case number > 69:
		fmt.Println("Grade: A")
	case number > 59:
		fmt.Println("Grade: B")
	default:
		fmt.Println("Bro, you failed 😭")
	}
}
