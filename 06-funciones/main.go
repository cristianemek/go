package main

import "fmt"

func main() {
	PrintOK("bien")
	ok := PrintOK("otro mensaje")
	fmt.Println(ok) //imprime el msg del return de la func
}

func PrintOK(msg string) string {
	fmt.Printf("OK - %s\n", msg)
	return msg
}
