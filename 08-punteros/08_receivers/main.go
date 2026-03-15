package main

import "fmt"

type Counter struct {
	number int
}

func (counter *Counter) Increment() {
	counter.number++
}

func (counter Counter) Value() int {
	return counter.number
}

func main() {
	counter := Counter{}
	counter.Increment()

	fmt.Println("Numero: ", counter.Value())
}
