package main

func mapIterator() {
	my_map := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	// Iterar sobre el mapa varias veces, el orden de los elementos no es garantizado, en cada iteración el orden puede cambiar
	for i := 0; i < 3; i++ {
		for key, value := range my_map {
			println("Key: ", key, "Value: ", value)
		}

	}
}
