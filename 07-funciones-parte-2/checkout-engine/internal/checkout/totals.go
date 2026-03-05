package checkout

func NewOrder(id, customer string) Order {
	return Order{
		ID:       id,
		Customer: customer,
		Items:    []Item{},
		Meta:     map[string]string{},
	}
}

// metodo para agregar item a la orden, convertir funcion en metodo (metodo es una funcion pero relacionada a un tipo de dato,en este caso Order),
// el receiver es el objeto al que se le asigna el metodo, en este caso order, se puede usar cualquier nombre pero por convención se usa la primera letra del tipo de dato
// se puede usar puntero para modificar el objeto original, si no se usa puntero se hace una copia del objeto y se modifica esa copia pero no el original
func (o *Order) AddItem(item Item) {
	o.Items = append(o.Items, item)
}

func (o *Order) RemoveItem(sku string) bool {
	for i := range o.Items {
		if o.Items[i].SKU == sku {
			//traer todos los elementos hasta la posicion i, separamos el slice en 2 partes, excluimos la posicion que se manda (i+1) y los ... para iterar 1 a 1 en el slice
			o.Items = append(o.Items[:i], o.Items[i+1:]...)
			return true
		}
	}
	return false
}

func CalcLineTotal(item Item) Money {
	return item.Price * Money(item.Qty)
}

func (o Order) CalcSubTotal() Money {
	var sum Money
	for _, item := range o.Items {
		sum += CalcLineTotal(item)
	}
	return sum
}

func (o Order) CalcTotalQuantity() int {
	var sum int
	for _, item := range o.Items {
		sum += item.Qty
	}
	return sum
}

func (o *Order) AddItems(items ...Item) {
	o.Items = append(o.Items, items...)
}

func (o Order) FindItem(sku string) (Item, bool) {
	for _, item := range o.Items {
		if item.SKU == sku {
			return item, true
		}
	}

	return Item{}, false
}

func (o Order) GetMeta(key string) (string, bool) {
	if o.Meta == nil {
		return "", false
	}
	value, exists := o.Meta[key]
	return value, exists
}

func (o Order) IndexOfItem(sku string) (int, bool) {
	for index, item := range o.Items {
		if item.SKU == sku {
			return index, true
		}
	}
	return -1, false
}

func (o Order) ApplyDiscounts(fns ...DiscountFn) Money {
	var discount Money
	for _, fn := range fns {
		discount += fn(o)
	}
	sub := o.CalcSubTotal()
	if discount > sub {
		return sub
	}

	return discount
}

func (o Order) Compute(tax TaxFn, bundle Money, ship ShippingFn, discount ...DiscountFn) (t Totals, err error) {
	//defer, se ejecuta al acabar la funcion Compute
	defer Track("Compute")() //la funcion track devuelve otra func,si no agrego () no se ejecuta la funcion que devuelve

	if err = ValidateOrder(o); err != nil {
		return Totals{}, err
	}

	t.Subtotal = o.CalcSubTotal()

	if bundle > 0 {
		t.Discount = bundle
	} else {
		t.Discount = o.ApplyDiscounts(discount...)

	}

	t.Tax = tax(o)
	t.Shipping = ship(o)
	t.Total = t.Subtotal - t.Discount + t.Tax + t.Shipping

	//se podria dejar el return vacio y go ya detecta que tipo de retorno tiene que dar, pero cuanto mas explícitos mejor
	return t, nil
}
