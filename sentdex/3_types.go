// Golang build int types are as follows
// bool

// string

// int  int8  int16  int32  int64
// uint uint8 uint16 uint32 uint64 uintptr

// byte // alias for uint8

// rune // alias for int32 represents a Unicode code point

// float32 float64

// complex64 complex128

package main

import ("fmt")

func add(x, y float64) float64 {
	return x+y
}
func subtract(x float64, y float64) float64 {
	return x-y
}
func multiply(x float64, y float64) float64 {
	return x*y
}
func divide(x float64, y float64) float64 {
	return x/y
}
func multiplyString(a,b string) (string,string) {
	return a,b
}
// func power(base float64, exponent float64) float64 {
// 	return base***exponent
// }
func main() {
	var num1, num2 float64= 2.1, 2
	// fmt.Println(add(2.1,2))
	fmt.Println(add(num1,num2))
	fmt.Println(multiplyString("abdulwahid","gul"))
}