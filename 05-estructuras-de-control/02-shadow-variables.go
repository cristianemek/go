package main

import "fmt"

func shadowVariables() {
	x := 10

	if x > 5 {
		fmt.Println(x)
		// con := estamos creando una varibale de 0 que solo existe en este bloque
		// x := 5
		// con = si que reasignamos valor y perdura fuera del bloque
		x = 5
		fmt.Println(x)
	}

	fmt.Println(x)
}
