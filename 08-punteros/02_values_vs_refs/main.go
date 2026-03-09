package main

import "fmt"

func incrementValue(number int) {
	number++
}

func mutateFirst(mySlice []int) { //slices y maps se mandan siempre como referencia
	mySlice[0] = 999
}

func mutateSecond(mySlice []int) { //slices y maps se mandan siempre como referencia
	mySlice[0] = 129
}

func push(s []int) []int {
	return append(s, 42)
}

func main() {
	x := 10
	incrementValue(x)

	fmt.Println("x: ", x)

	a := []int{1, 2, 3}
	mutateFirst(a)

	fmt.Println("a: ", a)

	b := push(a)

	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
	mutateSecond(b)
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)

}
