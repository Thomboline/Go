package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	// contact   contactInfo
	//OR JUST
	contactInfo
	grades []course
}

type contactInfo struct {
	email   string
	zipCode int
}

type course struct {
	name  string
	grade int
}

func main() {
	// One way to init: thomas := person1{"Thomas", "Christensen"}
	biology := course{name: "Biology", grade: 10}
	geography := course{name: "Geography", grade: 4}

	classes := []course{}
	classes = append(classes, biology, geography)

	// thomas := person{firstName: "Thomas", lastName: "Christensen", grades: classes}

	jim := person{
		firstName: "Jim",
		lastName:  "Hansen",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	jim.updateName("Thomas", "Christensen")
	jim.print()
}

func (p *person) updateName(fName string, lName string) {
	p.firstName = fName
	p.lastName = lName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
