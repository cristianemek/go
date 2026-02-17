package main

import (
	"fmt"
	"slices"
)

func slicesExample() {
	// //al no indicar ni el tamaño ni los ... el compilador infiere que es un slice y el tamaño es dinámico
	// var my_slice = []int{1, 2, 3}

	// fmt.Println(my_slice)
	// my_slice[1] = 10

	// fmt.Println(my_slice)
	// fmt.Println(my_slice[1])

	// //slice vacio
	// var empty_slice []int
	// //var empty_slice2 []int

	// //en los arrays el valor por defecto es 0, en los slices es nil==null
	// fmt.Println(empty_slice == nil) //true

	// //false, aunque ambos son nil, no son el mismo objeto en memoria, solo se puede comparar con nil
	// //fmt.Println(empty_slice2 == empty_slice)

	x := []int{1, 2, 3, 4, 5}
	y := []int{1, 2, 3, 4, 5}
	z := []int{1, 2, 3, 4, 5, 6}
	//s := []string{"a", "b", "c"}

	//fmt.Println(x == y) //error, no se pueden comparar dos slices, solo con nil

	fmt.Println(slices.Equal(x, y))
	fmt.Println(slices.Equal(x, z))
	//fmt.Println(slices.Equal(x, s)) //error, no se pueden comparar slices de diferentes tipos
}
