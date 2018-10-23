package main

import "fmt"

type car struct {
	gasPedal uint16
	brakePedal uint16
	steeringWheel int16
	topSpeedKhm float64
}

func main() {
	a_car :=  car{
		gasPedal: 22341,
		brakePedal: 0,
		steeringWheel: 12561,
		topSpeedKhm: 255.6}

	fmt.Println(a_car)
}