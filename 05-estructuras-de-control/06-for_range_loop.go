package main

func forRangeLoop() {
	evenNumbers := []int{2, 4, 6, 8, 10, 12}

	//indice y valor
	for index, value := range evenNumbers {
		println("Index: ", index, "Value: ", value)
	}

	//solo indice
	for index := range evenNumbers {
		println("Index: ", index)
	}

	//solo valor con guion bajo _ para ignorar el indice
	for _, value := range evenNumbers {
		println("Value: ", value)
	}
}
