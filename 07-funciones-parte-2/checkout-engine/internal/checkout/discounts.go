package checkout

import "strings"

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
