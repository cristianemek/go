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
		Price: 3000,
		Qty:   2,
	})

	//Validador
	PrintKV("Validador", ValidateOrder(order))

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

	PrintDivider()

	TryChangeCustomerByValue(order, "Cristian")

	PrintKV("Customer after TryChangeCustomerByValue", order.Customer)

	TryChangeCustomerByPointer(&order, "Cristian Updated")

	PrintKV("Customer after TryChangeCustomerByPointer", order.Customer)

	setCity(order, "Guadalajara")
	PrintKV("City", order.Meta["city"])

	PrintDivider()

	items := []Item{
		{SKU: "as-123", Name: "Teclado Membrana", Price: 3500, Qty: 1},
		{SKU: "Dd-133", Name: "Monitor Gaming", Price: 13500, Qty: 2},
	}

	AddItems(&order, items...)
	PrintKV("Cantidad Total: ", CalcTotalQuantity(order))
	PrintKV("Items: ", order.Items)

	PrintDivider()

	findItem, extraValueFind := FindItem(order, "Dd-133")

	Print2("Item encontrado", findItem, extraValueFind)

	getMeta, extraGetMeta := GetMeta(order, "city")

	Print2("Metadato encontrado", getMeta, extraGetMeta)
	IndexOfItemValue, INdexOfItemExtra := IndexOfItem(order, "as-123")
	Print2("Index encontrado", IndexOfItemValue, INdexOfItemExtra)

	PrintDivider()

	couponValue, couponError := ParseCoupon("SAVE10")

	Print2("Probando cupon: ", couponValue, couponError)

	PrintDivider()

	computeValue, computeError := Compute(order)
	Print2("Computar valores por nombre (TOTALES): ", computeValue, computeError)

	PrintDivider()

	PrintKV("Descuento", FlatDiscount(200)(order))
	th := ThresholdPercentDiscount(2000, 10)
	PrintKV("Descuento %: ", th(order))

	PrintDivider()

	//funcion anonima, la declaro y se la asigno a una varibale, y puedo llamarla a continuacion, util para casos puntuales y logica rapida
	cityDiscount := func(order Order) Money {
		city, _ := GetMeta(order, "city")
		if city == "Tijuana" {
			return 200
		}
		return 0
	}

	PrintKV("Descuento especial por ciudad: ", cityDiscount(order))
}
