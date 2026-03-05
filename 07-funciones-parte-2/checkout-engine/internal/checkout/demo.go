package checkout

func RunDemo() {
	PrintHeader("Hola Checkout Engine :)")
	order := NewOrder("123456789", "Cristian")
	order.AddItem(Item{
		SKU:   "ABC-123",
		Name:  "Teclado Mecánico",
		Price: 3500,
		Qty:   1,
	})

	order.AddItem(Item{
		SKU:   "ABD-133",
		Name:  "Monitor",
		Price: 13500,
		Qty:   2,
	})

	order.AddItem(Item{
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

	remove := order.RemoveItem("ABC-123")

	PrintKV("Item ABC-123 eliminado: ", remove)
	PrintKV("Items", len(order.Items))

	PrintDivider()

	sub := order.CalcSubTotal()
	qty := order.CalcTotalQuantity()

	PrintKV("Subtotal", StringUSD(sub))
	PrintKV("Total Quantity", StringUSD(qty))

	PrintDivider()

	TryChangeCustomerByValue(order, "Cristian")

	PrintKV("Customer after TryChangeCustomerByValue", order.Customer)

	TryChangeCustomerByPointer(&order, "Cristian Updated")

	PrintKV("Customer after TryChangeCustomerByPointer", order.Customer)

	setCity(&order, "Guadalajara")
	PrintKV("City", order.Meta["city"])

	//setear zona
	setZone(&order, "NATIONAL")

	PrintDivider()

	items := []Item{
		{SKU: "as-123", Name: "Teclado Membrana", Price: 3500, Qty: 1},
		{SKU: "Dd-133", Name: "Monitor Gaming", Price: 13500, Qty: 2},
	}

	order.AddItems(items...)
	PrintKV("Cantidad Total: ", order.CalcTotalQuantity())
	PrintKV("Items: ", order.Items)

	PrintDivider()

	findItem, extraValueFind := order.FindItem("Dd-133")

	Print2("Item encontrado", findItem, extraValueFind)

	getMeta, extraGetMeta := order.GetMeta("city")

	Print2("Metadato encontrado", getMeta, extraGetMeta)
	IndexOfItemValue, INdexOfItemExtra := order.IndexOfItem("as-123")
	Print2("Index encontrado", IndexOfItemValue, INdexOfItemExtra)

	PrintDivider()

	couponValue, couponError := ParseCoupon("SAVE10")

	Print2("Probando cupon: ", couponValue, couponError)

	PrintDivider()

	// computeValue, computeError := Compute(order)
	// Print2("Computar valores por nombre (TOTALES): ", computeValue, computeError)

	PrintDivider()

	PrintKV("Descuento", StringUSD(FlatDiscount(200)(order)))
	th := ThresholdPercentDiscount(2000, 10)
	PrintKV("Descuento %: ", StringUSD(th(order)))

	PrintDivider()

	//funcion anonima, la declaro y se la asigno a una varibale, y puedo llamarla a continuacion, util para casos puntuales y logica rapida
	cityDiscount := func(order Order) Money {
		city, _ := order.GetMeta("city")
		if city == "Tijuana" {
			return 200
		}
		return 0
	}

	PrintKV("Descuento especial por ciudad: ", cityDiscount(order))

	PrintDivider()

	discoutnKeyboard := MakeSKUDiscount("as-123", 500)
	discountHDMI := MakeSKUDiscount("Dd-133", 1000)

	PrintKV("Descuento por teclado: ", StringUSD(discoutnKeyboard(order)))
	PrintKV("Descuento por monitor: ", StringUSD(discountHDMI(order)))

	state, _ := order.GetMeta("city")
	zone, _ := order.GetMeta("zone")

	taxFn := NewTaxByState(state)
	shipFn := NewShippingByZone(zone)

	promo := CompositeDiscount{
		Name: "Promocion Febrero",
		Fns: []DiscountFn{
			FlatDiscount(100),
			ThresholdPercentDiscount(2000, 10),
			MakeSKUDiscount("as-123", 100),
		},
	}

	bundle := promo.Apply(order)
	PrintKV("DESCUENTO RECURSIVO", StringUSD(bundle))

	computeValue, computeError := order.Compute(taxFn, bundle, shipFn, FlatDiscount(5000), ThresholdPercentDiscount(2000, 10))
	Print2("Computar valores por nombre (TOTALES): ", computeValue, computeError)

	PrintKV("TOTAL: ", StringUSD(computeValue.Total))

}
