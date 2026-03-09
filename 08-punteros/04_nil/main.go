package main

import "fmt"

type User struct {
	name string
}

type Config struct {
	Host string
	Port int
}

func NewConfig() *Config {
	return &Config{Host: "localhost", Port: 8080}
}

func printName(user *User) {
	if user == nil {
		fmt.Println("User está vacío [nil]")
		return
	}
	fmt.Println(user.name)
}

func main() {
	var user *User
	printName(user)

	user = &User{name: "Cris"}
	printName(user)

	my_config := NewConfig()
	fmt.Println("configuracion: ", my_config.Host, my_config.Port)

}
