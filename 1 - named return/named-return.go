package main

import "fmt"

func calculosMatematicos(n1, n2 int) (suma, resta int) {
	suma = n1 + n2
	resta = n1 - n2
	return
}

func main() {
	_, resultadoResta := calculosMatematicos(10, 5)
	fmt.Println(resultadoResta)
}
