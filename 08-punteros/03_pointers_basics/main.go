package main

import "fmt"

func incrementValue(number *int) {
	*number++
}

type User struct {
	Name string
	Age  int
}

func main() {
	x := 10
	incrementValue(&x)

	fmt.Println("x: ", x)

	user := User{
		Name: "Cris",
		Age:  24,
	}

	birthday(&user)

	fmt.Println("User: ", user)

	a, b := 1, 2
	swap(&a, &b)

	fmt.Println("Swap values", a, b)

}

func birthday(user *User) {
	user.Age++
}

func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp

}
