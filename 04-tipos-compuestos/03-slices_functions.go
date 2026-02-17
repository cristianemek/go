package main

import "fmt"

func slicesFunctions() {
	// // 	my_slice := []int{}
	// // 	fmt.Println(my_slice == nil) //false, al indicar con []int{} el slice no es nil, aunque esté vacío, es un objeto en memoria
	// my_slice2 := []int{1, 2, 3}

	// //len
	// fmt.Println(len(my_slice2))

	// //apend
	// my_slice2 = append(my_slice2, 4, 5, 6, 7, 8, 9, 10)
	// fmt.Println(my_slice2)

	// //capacity,es la cantidad de elementos que el slice puede contener sin necesidad de reasignar memoria, cuando se supera la capacidad, el slice se redimensiona automáticamente, generalmente duplicando su capacidad actual para optimizar el rendimiento.
	// fmt.Println(cap(my_slice2))

	//crea un slice de enteros con longitud 5 y capacidad 5, los elementos se inicializan con el valor cero del tipo, en este caso 0
	// make_slice := make([]int, 5)

	//crea un slice de enteros con longitud 0 y capacidad 10, los elementos se inicializan con el valor cero del tipo, en este caso 0
	make_slice := make([]int, 0, 10)

	fmt.Println(make_slice)

	make_slice = append(make_slice, 10, 20, 30, 40, 50, 60, 70, 80, 90, 100)
	fmt.Println(make_slice)

	//aunque se indica la capacidad de 10, al superar esa capacidad el slice se redimensiona automáticamente, duplicando su capacidad inciall para optimizar el rendimiento.
	fmt.Println(len(make_slice), cap(make_slice))

	//vaciar slice, aunque el slice quede vacío, sigue existiendo en memoria, por lo que no es nil, solo se puede comparar con nil para saber si el slice ha sido inicializado o no

	//elimina todos los valores lo convierte a []
	// make_slice = make_slice[:0]
	// fmt.Println(make_slice)

	//clear convierte todos los valores del slice a su valor cero, en este caso 0
	clear(make_slice)
	fmt.Println(make_slice)

	fmt.Println(make_slice == nil) //false, el slice sigue existiendo en memoria aunque esté vacío
}
