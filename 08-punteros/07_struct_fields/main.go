package main

import "fmt"

type Address struct {
	City  string
	State string
}

type Profile struct {
	Bio string
}

type User struct {
	Name    string
	Addr    Address
	Profile *Profile //opcional o nulo
}

func main() {
	user := User{
		Name: "Ana",
		Addr: Address{
			City:  "Lugo",
			State: "Galicia",
		},
	}

	if user.Profile == nil {
		fmt.Println("Sin perfil")

	}

	user.Profile = &Profile{
		Bio: "Soy Cristian",
	}

	fmt.Println("Bio: ", user.Profile.Bio)
	fmt.Println(user)
}
