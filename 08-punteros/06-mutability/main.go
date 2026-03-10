package main

import "fmt"

type Order struct {
	Total int
}

func ApplyDiscount(order *Order, percent int) {
	order.Total = order.Total - (order.Total*percent)/100
}

func DiscountedTotal(order Order, percent int) Order {
	order.Total = order.Total - (order.Total*percent)/100
	return order
}

func main() {
	order := Order{Total: 1000}
	ApplyDiscount(&order, 20)
	fmt.Println(order)

	order2 := Order{
		Total: 1000,
	}
	discounted := DiscountedTotal(order2, 10)
	fmt.Println(order2.Total)
	fmt.Println(discounted)

}
