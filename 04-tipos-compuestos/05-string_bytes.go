package main

import "fmt"

func stringsBytes() {
	// var my_string string = "hola mundo"
	// var my_byte byte = my_string[0]

	// fmt.Println(my_byte)
	// fmt.Println(my_string[0]) //imprime el byte, no el caracter

	// var s2 string = my_string[0:4] //posicion 4 no se incluye
	// var s3 string = my_string[:7]  //desde el inicio hasta la posicion 4, sin incluirla

	// fmt.Println(s2)
	// fmt.Println(s3)

	//string a slices
	var s string = "hola mundo"

	//creo un slice de bytes a partir del string s, cada byte representa un caracter del string
	var s_slice []byte = []byte(s)
	var r_slice []rune = []rune(s)

	fmt.Println(s_slice)
	fmt.Println(r_slice)

}
