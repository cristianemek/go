package main

import "fmt"

func inputOutput() {
	var name string
	var age int

	//valores por defecto de los tipos de datos, en este caso name es una cadena vacía y age es 0
	fmt.Println(name, age)

	//dirección de memoria de las variables, se muestra la dirección de memoria donde se almacenan las variables name y age
	fmt.Printf("La dirección de memoria de name es: %p\n", &name)
	fmt.Printf("La dirección de memoria de age es: %p\n", &age)

	//ingresar datos por consola, se le pide al usuario que ingrese su nombre y edad, y se asignan a las variables name y age respectivamente

	fmt.Println("Ingresa tu nombre: ")
	fmt.Scan(&name)

	fmt.Println("Ingresa tu edad: ")
	//& es el operador que indica la dirección de memoria de la variable age, para que fmt.Scan pueda asignar el valor ingresado a esa variable
	fmt.Scan(&age)

	fmt.Printf("Hola %s, tienes %d años\n", name, age)
}
