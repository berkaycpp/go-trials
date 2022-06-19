package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	var berkay person
	berkay.firstName = "Berkay"
	berkay.lastName = "Sahin"
	berkay.contactInfo.email = "berkaysahin.c@gmail.com"
	berkay.contactInfo.zipCode = 12345

	someone := person{
		firstName: "asd",
		lastName:  "fgh",
		contactInfo: contactInfo{
			email:   "asd@fgh.com",
			zipCode: 67890,
		},
	}

	someone.updateName("xcv")
	someone.print()
}

func (p person) print() {
	fmt.Printf("\n%+v", p)
}

func (pointer2Person *person) updateName(newFirstName string) {
	(*pointer2Person).firstName = newFirstName
}
