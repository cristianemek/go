package checkout

func RunDemo() {
	PrintHeader("Hola Checkout Engine :)")
	order := NewOrder("123456789", "Cristian")
	AddItem(&order, Item{
		SKU:   "ABC-123",
		Name:  "Teclado Mecánico",
		Price: 3500,
		Qty:   1,
	})

	AddItem(&order, Item{
		SKU:   "ABD-133",
		Name:  "Monitor",
		Price: 13500,
		Qty:   2,
	})

	AddItem(&order, Item{
		SKU:   "ABD-123",
		Name:  "Monitor",
		Price: 1500,
		Qty:   2,
	})

	PrintKV("Order ID", order.ID)
	PrintKV("Customer", order.Customer)
	PrintKV("Items", len(order.Items))

	remove := RemoveItem(&order, "ABC-123")

	PrintKV("Item ABC-123 eliminado: ", remove)
	PrintKV("Items", len(order.Items))

	PrintDivider()

	sub := CalcSubTotal(order)
	qty := CalcTotalQuantity(order)

	PrintKV("Subtotal", sub)
	PrintKV("Total Quantity", qty)

}
