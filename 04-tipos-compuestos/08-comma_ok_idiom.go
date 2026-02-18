package main

import "fmt"

func commaOkIdiom() {
	my_map := map[string]int{
		"hola":  1,
		"mundo": 2,
	}

	//comma ok idiom, separas valor del map en 2, el ok es el booleano que nos dice si el valor existe o no
	value, ok := my_map["hAla"]

	fmt.Println(value, ok)

}
