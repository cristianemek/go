package main

import "fmt"

func mapsDelete() {
	my_map := map[string]int{
		"hola":  1,
		"mundo": 2,
	}

	fmt.Println(my_map)

	delete(my_map, "hola") //limpiar un elemento
	clear(my_map)          //limpias todos los elementos

	fmt.Println(my_map)
	fmt.Println(len(my_map))

}
