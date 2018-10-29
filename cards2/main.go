package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}
type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	vova := person{
		firstName: "Volodymyr",
		lastName:  "Paskiv",
		contact: contactInfo{
			email:   "mr.patriotuk@gmail.com",
			zipCode: 80600,
		},
	}
	vovaPointer := &vova //pointer
	vovaPointer.updateName("Volodya")
	vova.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPersone *person) updateName(n string) {
	(*pointerToPersone).firstName = n
}
