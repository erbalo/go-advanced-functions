package main

import "fmt"

func invierteValor(numero int) int {
	return numero * -1
}

func invierteValorPuntero(numero *int) {
	*numero = *numero * -1
}

func main() {
	numero := 20
	numeroInvertido := invierteValor(numero)
	fmt.Println(numeroInvertido)
	fmt.Println(numero)

	numero2 := 10
	fmt.Println(numero2)
	invierteValorPuntero(&numero2)
	fmt.Println(numero2)
}
