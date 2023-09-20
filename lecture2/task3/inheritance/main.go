package main

import "fmt"

type Person interface {
	Say() string
}

type person struct {
	FirstName string
	LastName  string
	Age       int
}

func (p person) Say() string {
	return "i am person"
}

type Student struct {
	person
	ID int64
}

func (s Student) Say() string {
	return "i am hungry"
}

type Employee struct {
	person
	Salary float64
}

func (e Employee) Say() string {
	return "i am tired"
}

func main() {
	p := person{"Alfarabi", "Agadilkhan", 19}
	s := Student{person: p, ID: 1}
	fmt.Print(s.Say())
}
