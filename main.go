package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	//david := person{firstName: "David", lastName: "Levy"}
	var david person
	fmt.Println(david)
}
