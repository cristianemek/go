package main

import "fmt"

// constanto no tipada se puede asignar a cualquier tipo compatible float64, float32..
// const Pi = 3.14159

// constanto tipada solo se puede asignar a su mismo tipo
const Pi float64 = 3.14159

func getTypedConstants() {
	var number float32 = float32(Pi)
	//float32 no se puede asignar a Pi porque es un float64, solo se puede asignar a Pi si se hace una conversión explícita a float32
	fmt.Println(number)
}
