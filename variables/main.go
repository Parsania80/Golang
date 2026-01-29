package main

import "fmt"

func main() {
	var age = 10
	fmt.Println("age is: ", age)
	var name = "mohammadreza"
	// fmt.Print("My name is ", name)
	_ = name

	s := "Test declearying variable"
	fmt.Println(s)

	car, cost := "BMW", 50000
	fmt.Println(car, cost)
}
