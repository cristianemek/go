package checkout

type TaxFn func(Order) Money

func NoTax(Order) Money {
	return 0
}

func IVA21(order Order) Money {
	sub := order.CalcSubTotal()
	return sub * 21 / 100
}

func NewTaxByState(state string) TaxFn {
	switch state {
	case "CDMX":
		return func(o Order) Money { return o.CalcSubTotal() * 16 / 100 }
	case "NL":
		return func(o Order) Money { return o.CalcSubTotal() * 16 / 100 }
	case "QRO":
		return func(o Order) Money { return o.CalcSubTotal() * 20 / 100 }
	case "GDL":
		return func(o Order) Money { return o.CalcSubTotal() * 14 / 100 }
	default:
		return NoTax
	}

}
