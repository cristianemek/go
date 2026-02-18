package main

import "fmt"

type Item struct {
	SKU   string
	Name  string
	Price int
	Qty   int
}

type Order struct {
	ID       string
	Customer string
	Items    []Item
	Meta     map[string]string
}

func main() {
	order := Order{
		ID:       "ORDER-1001",
		Customer: "Cristian",
		Items: []Item{
			{
				SKU:   "SKU-001",
				Name:  "Laptop",
				Price: 1200,
				Qty:   1,
			},
			{
				SKU:   "SKU-002",
				Name:  "Mouse",
				Price: 25,
				Qty:   2,
			},
		},
		Meta: map[string]string{
			"status": "completed",
			"city":   "Madrid",
			"source": "web",
		},
	}

	fmt.Println("Order ID: ", order.ID)
	fmt.Println("Customer: ", order.Customer)
	fmt.Println("Primer item: ", order.Items[0], "SKU:", order.Items[0].SKU)
	fmt.Println("Ciudad: ", order.Meta["city"])
	order.Meta["cupon"] = "DESCUENTO10"
	fmt.Println("Cupon aplicado:", order.Meta["cupon"])
	fmt.Println(order)
}
