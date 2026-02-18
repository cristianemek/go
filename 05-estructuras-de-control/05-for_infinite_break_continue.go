package main

import "fmt"

func forInfiniteBreakContinue() {

	for i := 0; i < 10; {
		fmt.Println("Iteración: ", i)
		if i%2 == 0 {
			i++
		} else {
			i += 2
		}
	}
}
