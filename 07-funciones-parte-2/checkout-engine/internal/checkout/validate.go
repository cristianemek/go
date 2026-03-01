package checkout

import (
	"errors"
	"fmt"
	"strings"
)

// sin usar puntero lo que se hace es como una copia de Order y al escribir en el modificamos esa copia pero no el original
func TryChangeCustomerByValue(o Order, name string) {
	// o.Customer = name
}

// aqui  al usar * pedimos el puntero, la direccion de memoria entonces ya no es una copia si no el objeto original q mandamos
func TryChangeCustomerByPointer(o *Order, name string) {
	o.Customer = name
}

// aqui aunque se pase order por valor, hay tipos de datos que aunque copiemos se pasan por referencia, como el map de meta, la copia de order recibe la referencia el puntero del map, se cambia la copia y el original
func setCity(o *Order, city string) {
	if o.Meta == nil {
		o.Meta = map[string]string{}
	}
	o.Meta["city"] = city // Map, Slice, func, pointer, chan
}

func setZone(o *Order, zone string) {
	if o.Meta == nil {
		o.Meta = map[string]string{}
	}
	o.Meta["zone"] = zone // Map, Slice, func, pointer, chan
}

func ValidateOrder(order Order) error {
	if order.ID == "" {
		return errors.New("El ID de la orden es obligatorio")
	}

	if order.Customer == "" {
		return errors.New("El cliente es obligatorio")
	}

	if len(order.Items) == 0 {
		return errors.New("La orden debe de tener almenos 1 elemento")
	}

	for i, item := range order.Items {
		if item.SKU == "" {
			return fmt.Errorf("Elemento[%d]: El SKU es obligatorio", i)
		}

		if item.Qty <= 0 {
			return fmt.Errorf("Item[%s]: su cantidad debe ser mayor a 0", item.SKU)
		}

		if item.Price < 0 {
			return fmt.Errorf("Item[%s]: precio debe ser mayor a 0", item.SKU)
		}
	}
	return nil
}

func ParseCoupon(code string) (Coupon, error) {
	coupon := strings.TrimSpace(strings.ToUpper(code))

	if coupon == "" {
		return Coupon{}, errors.New("Cupon vacio")
	}

	switch coupon {
	case "SAVE10":
		return Coupon{Code: coupon, Kind: "PERCENT", Val: 10}, nil
	case "LESS500":
		return Coupon{Code: coupon, Kind: "FLAT", Val: 500}, nil
	case "FREESHIP":
		return Coupon{Code: coupon, Kind: "FREESHIP", Val: 0}, nil
	default:
		return Coupon{}, fmt.Errorf("Cupon %q: es invalido", code)
	}
}
