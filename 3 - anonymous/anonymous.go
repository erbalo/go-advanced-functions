package main

import "fmt"

var algunaFuncion3 = func(text string) string {
	return fmt.Sprintf("Recbi string [%s]", text)
}

func main() {
	algunaFuncion := func(text string) string {
		return fmt.Sprintf("Recbi string [%s]", text)
	}

	algunaFuncion2 := func(text string) string {
		return fmt.Sprintf("Recbi string [%s]", text)
	}("Mario")

	fmt.Println(algunaFuncion("Erick"))

	fmt.Println(algunaFuncion2)

	fmt.Println(("scope"))
}
