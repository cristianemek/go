package main

import "fmt"

func anonymousStruct() {
	var person struct {
		name string
		age  int
		pet  string
	}

	person.name = "Cristian"
	person.age = 30
	person.pet = "dog"

	pet := struct {
		name string
		age  int
	}{
		name: "Buddy",
		age:  5,
	}

	fmt.Println(person, pet)
}
