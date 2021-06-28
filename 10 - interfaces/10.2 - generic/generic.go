package main

import "fmt"

func generica(inter interface{}) {
	fmt.Println(inter)
}

func main() {
	generica("Strings")
	generica(true)
	generica(1)
	generica(1.353)

	mapa := map[interface{}]interface{}{
		"12412":   true,
		true:      1,
		216712.21: "Erick",
		18:        "Otra cosa",
	}

	fmt.Println(mapa)
}
