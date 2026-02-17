package main

import "fmt"

func maps() {

	// var nilMap map[string]int

	//nilMap["key"] = 1
	//esto da un panic porque el mapa no ha sido inicializado, es nil

	//con {} se inicializa el mapa ya no es nil, se puede usar
	totalWins := map[string]int{}
	fmt.Println(totalWins)

	totalWins["Cristian"] = 1
	totalWins["Maria"] = 2
	totalWins["Cristian"] = 3

	fmt.Println(totalWins)
	fmt.Println(totalWins["Cristian"])

}
