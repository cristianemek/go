package main

import "fmt"

func main() {
	// PrintOK("bien")
	// ok := PrintOK("otro mensaje")
	// fmt.Println(ok) //imprime el msg del return de la func

	// demoDeferLIFO()
	fmt.Println(factorial(5))
	fmt.Println(sum([]int{1, 2, 3, 4}))
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

func factorial(number int) int {
	if number < 0 {
		panic("Numero debe ser mayor o igual a cero (0)")
	}

	if number == 0 {
		return 1
	}

	return number * factorial(number-1)
}

func sum(nums []int) int {
	//caso base:

	if len(nums) == 0 {
		return 0
	}

	return nums[0] + sum(nums[1:])
}
