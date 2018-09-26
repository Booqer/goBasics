package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type doubleZero struct {
	person
	first         string
	licenseToKill bool
}

func main() {
	p1 := doubleZero{
		person: person{
			first: "James",
			last:  "Bond",
			age:   20,
		},
		first:         "007",
		licenseToKill: true,
	}
	fmt.Println(p1)
	//  fmt.Printf("Name: ", %v, %v, "Age: ", %v, "Code Name: ", %v, "License To Kill: ", %v, p1.person.first, p1.person.last, p1.person.age, p1.first, p1.licenseToKill)
}
