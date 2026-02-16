package main

import "fmt"

//no se puede declarar una short variable declaration fuera de una función, ya que solo se permite dentro de funciones, por lo que se debe usar la declaración de variable tradicional con var
//name := "Cristian"

func declaration() {

	//declaracion de variable sin asignar un valor, por lo que se le asigna el valor cero del tipo int, que es 0
	var age int

	//short variable declaration, se asigna el valor "Cristian" a la variable name, y el tipo de la variable se infiere automáticamente como string
	name := "Cristian"

	fmt.Printf("My name is %s and I am %d years old\n", name, age)
}
