package main

import "fmt"

func closure() func() {
	text := "Dentro de la funcion closure"

	funcion1 := func() {
		fmt.Println(text)
	}

	return funcion1
}

func main() {
	text := "Dentro de la funcion main"
	fmt.Println(text)

	funcionNueva := closure()
	funcionNueva()
}
