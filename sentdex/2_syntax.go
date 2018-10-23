package main

import ("fmt"
		"math"
		"math/rand"
		"time")


func getSqrt(value float64) float64 {
	sendThis := math.Sqrt(value)
	return sendThis
}

func getRandom(value int) int {
	sendThis := rand.Intn(value)
	return sendThis
}

func main() {
	fmt.Println("The Squre root of 12329 is", getSqrt(12329))
	fmt.Println("The Random value between 1-100 is",getRandom(100))
	fmt.Println("The current time is ", time.Now())
}