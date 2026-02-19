package checkout

func NewOrder(id, customer string) Order {
	return Order{
		ID:       id,
		Customer: customer,
		Items:    []Item{},
		Meta:     map[string]string{},
	}
}

// el *
func AddItem(order *Order, item Item) {
	order.Items = append(order.Items, item)
}

func RemoveItem(order *Order, sku string) bool {
	for i := range order.Items {
		if order.Items[i].SKU == sku {
			//traer todos los elementos hasta la posicion i, separamos el slice en 2 partes, excluimos la posicion que se manda (i+1) y los ... para iterar 1 a 1 en el slice
			order.Items = append(order.Items[:i], order.Items[i+1:]...)
			return true
		}
	}
	return false
}

func CalcLineTotal(item Item) Money {
	return item.Price * Money(item.Qty)
}

func CalcSubTotal(order Order) Money {
	var sum Money
	for _, item := range order.Items {
		sum += CalcLineTotal(item)
	}
	return sum
}

func CalcTotalQuantity(order Order) int {
	var sum int
	for _, item := range order.Items {
		sum += item.Qty
	}
	return sum
}

func AddItems(order *Order, items ...Item) {
	order.Items = append(order.Items, items...)
}

func FindItem(order Order, sku string) (Item, bool) {
	for _, item := range order.Items {
		if item.SKU == sku {
			return item, true
		}
	}

	return Item{}, false
}

func GetMeta(order Order, key string) (string, bool) {
	if order.Meta == nil {
		return "", false
	}
	value, exists := order.Meta[key]
	return value, exists
}

func IndexOfItem(order Order, sku string) (int, bool) {
	for index, item := range order.Items {
		if item.SKU == sku {
			return index, true
		}
	}
	return -1, false
}
