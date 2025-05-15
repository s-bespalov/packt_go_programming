package main

import "fmt"

type Speaker interface {
	Speak() string
}

type person struct {
	name      string
	age       int
	isMarried bool
}

func main() {
	p := person{name: "Caylin", age: 44, isMarried: false}
	fmt.Println(p)
	fmt.Println(p.Speak())
}

func (p person) String() string {
	return fmt.Sprintf("%v (%v years old). \nMarried status: %v ", p.name, p.age, p.isMarried)
}

func (p person) Speak() string {
	return "Hi my name is " + p.name
}
