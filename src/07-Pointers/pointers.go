package main

import (
	"fmt"
	"reflect"
	"strings"
)

func changeName(n *string) {
	*n = strings.ToUpper(*n)
}

type Coordinates struct {
	X, Y float64
}

func main() {
	var name string = "Zoyah"
	var namePointer *string = &name

	fmt.Println(name, namePointer)
	fmt.Println(reflect.TypeOf(name), reflect.TypeOf(namePointer))
	fmt.Println(*namePointer)

	me := "tiskae"
	changeName(&me)
	fmt.Println(me)

	c := Coordinates{X: 7.0000, Y: 4.0000}
	cAddress := &c

	cAddress.X = 7.5000

	fmt.Println(c)
}
