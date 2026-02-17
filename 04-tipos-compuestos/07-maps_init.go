package main

import "fmt"

func mapsInit() {

	//Map literal

	// teams := map[string][]string{
	// 	"Team A": {"Player 1", "Player 2", "Player 3"},
	// 	"Team B": {"Player 4", "Player 5", "Player 6"},
	// }

	// fmt.Println(teams)
	// fmt.Println(teams["Team A"])
	// fmt.Println(teams["Team A"][0])

	//pueden ser llaves de un mapa solo los tipos comparables (==,!=) , no se pueden usar slices, maps o funciones como llaves de un mapa
	ages := make(map[int][]string, 10)

	fmt.Println(ages)
}
