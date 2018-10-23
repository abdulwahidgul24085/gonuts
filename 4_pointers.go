package main

import "fmt"

func main() {
	x := 15
	a := &x
	fmt.Println("The mem address for a",a)
	fmt.Println("Read value at mem address for a ",*a)
	*a = 100
	fmt.Println("The value of the x is ", x)
	fmt.Println("The memory address of a", &a)
}