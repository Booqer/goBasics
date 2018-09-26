package main

import "fmt"

type person struct {
	name string
	age  int
}

type doubleZero struct {
	person
	licenseToKill bool
}

func (p person) Greeting() {
	fmt.Println("I'm just a regular person.")
}

func (dz doubleZero) Greeting() {
	fmt.Println("Miss Moneypenny, so good to see you.")
}

func main() {
	p1 := person{
		name: "Ian Flemming",
		age:  44,
	}

	p2 := doubleZero{
		person: person{
			name: "James Bond",
			age:  23,
		},
		licenseToKill: true,
	}

	p1.Greeting()
	p2.Greeting()
	p2.person.Greeting()
}
