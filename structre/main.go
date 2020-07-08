package main

import "fmt"

type contactInfo struct {
	zipCode string
	emai    string
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	// alex := person{firstName: "alex", lastName: "anderson"}
	// fmt.Println(alex)

	alex := person{
		lastName:  "Anderson",
		firstName: "Alex",
		contactInfo: contactInfo{
			zipCode: "123",
			emai:    "asdfa@esd.com",
		},
	}

	// alexPointer := &alex
	// alexPointer.updateName("alexia")
	alex.updateName("alexia")
	alex.print()
}

func (pointerToPerson *person) updateName(newName string) {
	(*pointerToPerson).firstName = newName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
