package main

import "fmt"

func add(x,y int) int {
	return x+y
}

func Location(city string) (string, string) {
	var region, country string

	switch city {
	case "Karachi", "KHI", "khi":
		region, country = "Sindh", "Pakistan"
	case "Lahore", "LHR", "lhr":
		region, country = "Punjab", "Pakistan"
	case "Quetta" , "QTA", "qta":
		region, country = "Balochistan", "Pakistan"
	default:
		region, country = "Kashimir", "Pakistan"
	}

	return region, country
}

type Artist struct {
	Name, Genre string
	Song int
}

// if we pass the value with out the *, then a copy of the struct is passed and this would lead to creating a copy, and not changing the source value.
func newRelease(a *Artist) int {
	a.Song++
	return a.Song
}

func main() {
	var (
			name string = "abdulwahid gul"
			age int = 27
			location string = "Pakistan"
			distance float64 = 1.22
			attendence bool = true
			genre string = "Sufi Music"
		)

	hobby := "Programming not"
	year := 2017

	company := func() {
		fmt.Println("\nfunction called")
	}
	fmt.Println(name,age,distance,attendence,location,nil)
	fmt.Printf("Hobbies include: %s since year %d",hobby,year)
	company()

	// Constants can only be character, string, boolean, or numeric values and cannot be declared using the := syntax
	const Pi = 3.14

	fmt.Println(Pi)
	print(add(age,year))

	region, country := Location("Karachi")
	fmt.Printf("\n%s lives in %s, %s", name,region, country)

	// we need to set the pointers for not creating a copy, and passing the value by refrence 
	me := &Artist{name,genre,1244}

	fmt.Printf("\n%s released his %dth Song",me.Name,newRelease(me))
	fmt.Printf("\n%s has a total of %d Songs", me.Name, me.Song)

	you :=&Artist{"hammad","Rap",1}
	fmt.Printf("\n%s released his %dth Song",you.Name,newRelease(you))
	fmt.Printf("\n%s has a total of %d Songs", you.Name, you.Song)	
}



