package main

import (
	"fmt"
	"main/src/05-Toolkit/utils"
)

func calculateImportantData() int {
	totalValue := utils.Add(1, 2, 3, 4, 5, 6)
	return totalValue
}

func main() {
	fmt.Println("Packages!")
	total := calculateImportantData()
	fmt.Println(total)
}
