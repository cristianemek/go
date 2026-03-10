package main

import "fmt"

type User struct {
	Name     string
	Age      int
	IsActive bool
}

type UpdateUser struct {
	Name     *string
	Age      *int
	IsActive *bool
}

func Applu(user *User, update UpdateUser) {
	if update.Name != nil {
		user.Name = *update.Name
	}

	if update.Age != nil {
		user.Age = *update.Age
	}

	if update.Age != nil {
		user.Age = *update.Age
	}
}

func main() {
	user := User{
		Name:     "Ana",
		Age:      20,
		IsActive: true,
	}

	fmt.Println(user)

	newName := ""
	zeroAge := 0
	inactive := false

	Applu(&user, UpdateUser{
		Name:     &newName,
		Age:      &zeroAge,
		IsActive: &inactive,
	})

	fmt.Println(user)

}
