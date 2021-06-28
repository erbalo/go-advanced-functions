package main

import "fmt"

type usuario struct {
	name string
	age  uint8
}

func (u usuario) esAdulto() bool {
	return u.age >= 18
}

func (u *usuario) hacerCumpleanos() {
	u.age++
}

func main() {
	usuario1 := usuario{"mario", 18}
	fmt.Println(usuario1)
	fmt.Println(usuario1.esAdulto())

	fmt.Println("---------")

	usuario2 := usuario{"gil", 17}
	fmt.Println(usuario2)
	fmt.Println(usuario2.esAdulto())
	usuario2.hacerCumpleanos()
	fmt.Println(usuario2)
	fmt.Println(usuario2.esAdulto())
}
