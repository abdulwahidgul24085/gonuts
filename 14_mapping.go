package main

import "fmt"

func main() {
	Grades := make(map[string]float32) // to make it have values, we use the make()
	
	Grades["Abdulwahid"] = 42
	Grades["Ali"] = 9
	Grades["suleman"] = 99

	fmt.Println(Grades)

	AbdulwahidGrades := Grades["Abdulwahid"]
	fmt.Println(AbdulwahidGrades)

	delete(Grades,"Abdulwahid")

	for k, v := range Grades {
		fmt.Println(k, ":",v)
	}
}