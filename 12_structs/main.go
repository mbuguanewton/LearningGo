package main

import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	// age       int

	// alternative cleaner way
	firstName, lastName, city, gender string
	age                               int
}

// Greeting method (value reciever)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " " + "and i am " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer reciever)

func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer reciever)
func (p *Person) getMarried(spouseLastname string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastname
	}
}

func main() {
	// Init person using struct
	person1 := Person{firstName: "Samantha", lastName: "Smith", city: "Boston", gender: "f", age: 25}

	// Alternative
	person2 := Person{"Bob", "Johnson", "Boston", "m", 35}
	fmt.Println(person1, person2)

	// get single field
	fmt.Println(person1.firstName)

	// print output from the value reciever
	person1.hasBirthday()
	person1.getMarried("McAdams")
	person2.getMarried("Thompson")
	fmt.Println(person1.greet())
	fmt.Println(person2.greet())

}
