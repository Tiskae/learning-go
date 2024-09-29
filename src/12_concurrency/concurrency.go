package main

import (
	"fmt"
	"sync"
	"time"
)

var waitgroup sync.WaitGroup

func handlePanic() {
	if r := recover(); r != nil {
		fmt.Println("Something went wrong but we are gonna be all fine", r)
	}
}

func printStuff() {
	defer waitgroup.Done()
	defer handlePanic()

	for i := 0; i < 3; i++ {
		fmt.Println(i)
		time.Sleep(time.Millisecond * 300)

		if i == 2 {
			panic("PANIC!")
		}
	}
}

func main() {
	waitgroup.Add(1)
	go printStuff()
	waitgroup.Wait()
}
