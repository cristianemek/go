package main

import "fmt"

func sliceOfSlice() {

	my_string_slice := []string{"a", "b", "c", "d", "e"}
	y := my_string_slice[0:2] //posicion 2 no se incluye
	z := my_string_slice[2:5]
	d := my_string_slice[2:4]

	//esto es una copia de la direccion de memoria, si modificamos el slice, se modifica el original
	// e := my_string_slice[:]

	e := make([]string, 5)
	copy_e := copy(e, my_string_slice)

	fmt.Println(copy_e)

	e[0] = "x"

	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(d)
	fmt.Println(e)

	fmt.Println(my_string_slice)
}
