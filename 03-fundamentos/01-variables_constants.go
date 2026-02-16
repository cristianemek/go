package main

import "fmt"

const App = "Curso Go" //app=interno, App=exportable, esto es una convención del lenguaje
const maxUsers = 1000

func constants() {
	var name string = "Cristian"
	last_name := "Rodríguez"

	fmt.Printf("Hola, %s %s %T\n", name, last_name, last_name)
	fmt.Println("App:", App, "Max Users:", maxUsers)
}
