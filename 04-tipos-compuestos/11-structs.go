package main

import "fmt"

func structsExample() {
	type person struct {
		name string
		age  int
		pet  string
	}

	// var cristian person
	// fernando := person{}

	cristian := person{
		"Cristian",
		30,
		"dog",
	}

	fernando := person{
		age:  25,
		pet:  "cat",
		name: "Fernando",
	}

	cristian.name = "Cristiann"
	fmt.Println(cristian.name, fernando.name)

}
