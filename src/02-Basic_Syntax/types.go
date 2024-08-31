package main

import (
	"fmt"
	"reflect"
)

func types() {
	name := "Rahmat"
	fmt.Println(reflect.TypeOf(name))

	var gpa float64 = 4.55
	var tnu int = 23
	var tcp = gpa * float64(tnu)

	fmt.Printf("Total cummulative point: %f", tcp)
}
