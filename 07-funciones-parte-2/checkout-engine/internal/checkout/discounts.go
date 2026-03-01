package checkout

import "strings"

type DiscountFn func(Order) Money //esperamos que entre para el tipo discountFn una func que recibe un order y regresa Money

type Coupon struct {
	Code string
	Kind string
	Val  int
}

type CompositeDiscount struct {
	Name string
	Fns  []DiscountFn
}

func ApplyCouponCodes(order *Order, codes ...string) {
	if order.Meta == nil {
		order.Meta = map[string]string{}
	}

	order.Meta["coupons"] = joinCoupons(codes)

}

func joinCoupons(coupons []string) string {
	if len(coupons) == 0 {
		return ""
	}

	var out strings.Builder
	out.WriteString(coupons[0])
	for i := 1; i < len(coupons); i++ {
		out.WriteString("," + coupons[i])
	}
	return out.String()
}

func FlatDiscount(amount Money) DiscountFn {
	return func(o Order) Money {
		return amount
	}
}

func ThresholdPercentDiscount(min Money, percent int) DiscountFn {
	return func(order Order) Money {
		sub := order.CalcSubTotal()
		if sub < min {
			return 0
		}
		return sub * Money(percent) / 100
	}
}

// Closure: función que regresa otra función, y esa función tiene acceso a las variables de la función padre
// lo potente de devolver una funcion es que no realizo ningun calculo aqui solo creo la funcion configurada,
// cuando se llame en el codigo es cuando se calcula
func MakeSKUDiscount(sku string, amount Money) DiscountFn {
	return func(o Order) Money {
		_, ok := o.FindItem(sku)
		if !ok {
			return 0
		}
		return amount
	}
}

func ApplyDiscountsRecursive(order Order, fns []DiscountFn) Money {
	if len(fns) == 0 {
		return 0
	}

	return fns[0](order) + ApplyDiscountsRecursive(order, fns[1:])
}

func (composite CompositeDiscount) Apply(order Order) Money {
	return ApplyDiscountsRecursive(order, composite.Fns)
}
