package main

import "fmt"

// Define person struct
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

func main() {
	// Init person using struct
	person1 := Person{firstName: "Jesxion", lastName: "Zhang", city: "Zhengzhou", gender: "m", age: 32}

	// Alternative
	// person1 := Person{"Jesxion", "Zhang", "Zhengzhou", "m", 32}

	fmt.Println(person1.firstName)
	person1.age++
	fmt.Println(person1)
}
