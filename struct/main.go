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
	/**
		//gabriel := person{firstName: "Gabriel", lastName: "Galvão"}
		//fmt.Println(gabriel)

		//var gabriel person
		//
		//gabriel.firstName = "Gabriel"
		//gabriel.lastName = "Galvão"
		//
		//fmt.Println(gabriel)
		//fmt.Printf("%+v", gabriel)
	**/

	frida := person{
		firstName: "Frida",
		lastName:  "Chagas",
		contactInfo: contactInfo{
			email:   "fridachagas@email.com",
			zipCode: 12345,
		},
	}

	//fridaPointer := &frida
	frida.updateName("Galvão")
	frida.print()
}

func (pointerToPerson *person) updateName(newLastName string) {
	(*pointerToPerson).lastName = newLastName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
