package main

import "fmt"

func switchCondition() {
loop:
	for i := 0; i < 10; i++ {
		//switch size := len(examples); size {}
		switch i {
		case 0, 2, 4, 6, 8:
			fmt.Println(i, "es un número par")
		case 3:
			fmt.Println(i, "es un número impar, pero es el 3")
		case 7:
			fmt.Println("Aqui termina el programa")
			break loop
		default:
			fmt.Println(i, "valor por defecto")

		}
	}
}
