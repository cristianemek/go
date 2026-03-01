package checkout

import "strings"

type DiscountFn func(Order) Money //esperamos que entre para el tipo discountFn una func que recibe un order y regresa Money

type Coupon struct {
	Code string
	Kind string
	Val  int
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
		sub := CalcSubTotal(order)
		if sub < min {
			return 0
		}
		return sub * Money(percent) / 100
	}
}
