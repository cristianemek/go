package checkout

// sin usar puntero lo que se hace es como una copia de Order y al escribir en el modificamos esa copia pero no el original
func TryChangeCustomerByValue(o Order, name string) {
	o.Customer = name
}

// aqui  al usar * pedimos el puntero, la direccion de memoria entonces ya no es una copia si no el objeto original q mandamos
func TryChangeCustomerByPointer(o *Order, name string) {
	o.Customer = name
}

// aqui aunque se pase order por valor, hay tipos de datos que aunque copiemos se pasan por referencia, como el map de meta, la copia de order recibe la referencia el puntero del map, se cambia la copia y el original
func setCity(o Order, city string) {
	if o.Meta == nil {
		o.Meta = map[string]string{}
	}
	o.Meta["city"] = city // Map, Slice, func, pointer, chan
}
