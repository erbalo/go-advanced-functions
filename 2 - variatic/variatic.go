package main

import "fmt"

func suma(numeros ...int) int {
	total := 0

	for _, numero := range numeros {
		total += numero
	}

	return total
}

func escribrirTotal(texto string, numeros ...int) {
	total := suma(1, 2, 4, 4, 5, 5, 6, 6)
	fmt.Printf("Text %s: %d", texto, total)
}

func main() {
	escribrirTotal("Hola mundo!", 1, 2, 3, 4, 5)
}
