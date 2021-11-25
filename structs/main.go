package main

import (
	"fmt"
)

type ContactInfo struct {
	email string
	zip   int
}

type Person struct {
	firstName   string
	lastName    string
	contactInfo ContactInfo
}

func (person Person) getFullName() string {
	return fmt.Sprintf("%v %v", person.firstName, person.lastName)
}

func (person Person) getContactInfo() string {
	return fmt.Sprintf("%v %v", person.contactInfo.email, person.contactInfo.zip)
}

func (person *Person) updateName() {
	person.firstName = "evan2"
}

func (person *Person) updateLastName() {
	(*person).lastName = "fleming2"
}

// &value   -> turn value into pointer
// *pointer -> turn pointer into value
func main() {
	person := Person{
		firstName: "evan",
		lastName:  "fleming",
		contactInfo: ContactInfo{
			email: "evan.fleming@dev.local",
			zip:   90210,
		},
	}

	person.updateName()

	personPointer := &person
	personPointer.updateLastName()

	fmt.Println(person)
	fmt.Printf("%+v", person)
	fmt.Println(person.getFullName())
	fmt.Println(person.getContactInfo())
}

// type Internal struct {
// 	value int
// }
// type Test struct {
// 	internal *Internal
// }

// func (test Test) updateInternal() {
// 	test.internal.value = 7
// }

// func main() {
// 	internal := Internal{
// 		value: 5,
// 	}
// 	test := Test{
// 		internal: &internal,
// 	}

// 	fmt.Println(test.internal.value)
// }
