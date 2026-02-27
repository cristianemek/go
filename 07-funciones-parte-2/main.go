package main

import "fmt"

func main() {
	PrintOK("bien")
	ok := PrintOK("otro mensaje")
	fmt.Println(ok) //imprime el msg del return de la func

	demoDeferLIFO()

}

func PrintOK(msg string) string {
	fmt.Printf("OK - %s\n", msg)
	return msg

	//LIFO - Last In First Out
}

// defer se ejecutan al final de la funcion, pero se guardan en una pila, entonces se ejecutan en orden inverso a como se declararon, el ultimo defer declarado es el primero en ejecutarse
func demoDeferLIFO() {
	x := 10
	defer fmt.Println("defer 1", x) //se ejecuta al final de la funcion, pero se guarda el valor de x en el momento que se declara, entonces imprime 10
	x = 99
	defer PrintOK("primero")
	defer PrintOK("segundo")
	PrintOK("tercero")
}
