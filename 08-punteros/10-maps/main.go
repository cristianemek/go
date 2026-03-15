package main

import "fmt"

func put(my_map map[string]int, key string, value int) {
	my_map[key] = value
}

func main() {
	var my_map map[string]int
	fmt.Println("leer nil map:", my_map["x"]) //  my_map == nil //true

	my_map = make(map[string]int)
	fmt.Println("Mapa vacío:", my_map == nil) //false
	put(my_map, "a", 10)
	put(my_map, "b", 20)
	fmt.Println("Mapa: ", my_map)
}
