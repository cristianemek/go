package main

import "fmt"

func main() {
	x := 10
	p := &x //accedemos a la direccion de x

	fmt.Println("x = ", x)
	fmt.Println("p = ", p)

	fmt.Println("*p = ", *p) //acceder al valor de la direccion

	*p = 99
	fmt.Println("*p = ", *p)
	fmt.Println("x = ", x)
}
