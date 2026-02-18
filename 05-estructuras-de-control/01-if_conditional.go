package main

import (
	"fmt"
	"math/rand"
)

func ifConditional() {
	// should_buy := true
	// debts := true

	// if should_buy {
	// 	if debts {
	// 		fmt.Println("No puedes comprar porque tienes deudas")
	// 	} else {
	// 		fmt.Println("Puedes comprar")
	// 	}
	// } else {
	// 	fmt.Println("No puedes")
	// }

	//enteros del 0 - 10
	if n := rand.Intn(10); n == 0 {
		fmt.Println("Numero 0")
	} else if n > 5 {
		fmt.Println("Número > 5", n)
	} else {
		fmt.Println("Numero ideal: ", n)
	}

}
