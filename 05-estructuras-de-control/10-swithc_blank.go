package main

import "fmt"

func switchBlank() {
	a := 4

	switch {
	case a == 2:
		fmt.Println("a es igual a 2")
	case a == 4:
		fmt.Println("a es igual a 4")
	default:
		fmt.Println("a no es ni 2 ni 4")
	}
}
