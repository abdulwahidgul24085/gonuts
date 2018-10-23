package main

import "fmt"
const sixteenBitMax float64 = 65535
const kmhMultiple float64 = 1.60934

type car struct {
	gasPedal uint16
	brakePedal uint16
	steeringWheel int16
	topSpeedKhm float64
}

// value receiver for a struct
func (c car) kmph() float64 {
	return float64(c.gasPedal) * (c.topSpeedKhm / sixteenBitMax)
}

// value receiver for a struct
func (c car) mph() float64 {
	return float64(c.gasPedal) * (c.topSpeedKhm / sixteenBitMax/kmhMultiple)
}

// pointer receiver. This modifies the internal values
func (c *car) newTopSpeed(topSpeed float64) {
	c.topSpeedKhm = topSpeed
}
func main() {
	a_car :=  car{
		gasPedal: 65530,
		brakePedal: 0,
		steeringWheel: 12561,
		topSpeedKhm: 225}
	
	fmt.Println(a_car)
	fmt.Println(a_car.kmph())
	fmt.Println(a_car.mph())
	a_car.newTopSpeed(500)
	fmt.Println(a_car.kmph())
	fmt.Println(a_car.mph())
}