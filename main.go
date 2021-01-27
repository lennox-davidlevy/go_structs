package main

import "fmt"

type contactInfo struct {
	email string
	zip   int
}
type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	david := person{
		firstName: "David",
		lastName:  "Levy",
		contactInfo: contactInfo{
			email: "blah",
			zip:   11120,
		},
	}
	davidPointer := &david
	davidPointer.updateName("Davey")
	david.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
func (p person) print() {
	fmt.Printf("%+v", p)
}
