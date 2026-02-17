package main

import "fmt"

func arrays() {

	// array literal
	var number_list = [3]int{1, 2, 3}

	//array simplificado, el compilador infiere el tamaño del array
	var number_list2 = [...]int{1, 2, 3}

	fmt.Println(number_list)
	fmt.Println(len(number_list2))

}
