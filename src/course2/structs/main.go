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
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// fmt.Println(alex.firstName + " " + alex.lastName)
	// fmt.Println(alex)
	var alex person
	alex.firstName = "charu"
	alex.lastName = "sharma"
	alex.contactInfo = contactInfo{email: "kt@abc.in", zipCode: 302012}
	fmt.Println(alex)
	fmt.Printf("%+v", alex)
	kt := person{
		firstName: "karvy",
		lastName:  "joel",
		contactInfo: contactInfo{
			email:   "kt@gm.com",
			zipCode: 11007,
		},
	}

	kt.updateName("parvati")
	kt.print()

}

func (p person) print() {
	fmt.Println()
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
