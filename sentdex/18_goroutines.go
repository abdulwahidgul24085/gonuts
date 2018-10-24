package main

import (
	"fmt"
	"time"
	"sync"
)

// weight group
// This waits for the go routine to complete
var wg sync.WaitGroup

func cleanup(){
	defer wg.Done()
	if r:= recover(); r != nil {
		fmt.Println("Recovered in cleanup", r)
	}
}

func say(s string) {
	// The defer statement will be evalutated immediately, but wont run until the end, so you can, and usually will, put the defer statement early in the function
	defer cleanup()
	for i:= 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 100)
		if i == 2 {
			panic("This does not work bro")
		}
	}
}

func foo() {
	defer fmt.Println("Done")
	defer fmt.Println("Are we Done?")
	fmt.Println("This is doing something, Like seriously")

}

func main() {
	// This add's wait to the routine
	wg.Add(1)
	go say("abdulwahid")
	wg.Add(1)
	go say("gul")
	wg.Wait()

	wg.Add(1)
	go say("Boy")
	wg.Wait()

	foo()
	// Time delay is introduced coz the code will not stop for the concurrent function to finish
	// time.Sleep(time.Second)
}