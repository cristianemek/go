package main

import "fmt"

const tax = 0.21

func main() {
	var clientName string
	var clientEmail string
	var subtotal float64

	fmt.Println("Ticket por consola")
	fmt.Print("Nombre del cliente: ")
	fmt.Scanln(&clientName)

	fmt.Print("Email del cliente: ")
	fmt.Scanln(&clientEmail)

	fmt.Print("Subtotal: ")
	fmt.Scanln(&subtotal)

	fmt.Println("***********************")

	fmt.Printf("Subtotal: %.2f\n", subtotal)
	fmt.Printf("Impuestos: %.2f\n", subtotal*tax)
	fmt.Printf("Total: %.2f\n", subtotal*(1+tax))

	fmt.Println("***********************")

}
