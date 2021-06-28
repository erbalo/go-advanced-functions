package main

import "fmt"

var numero int

func init() {
	fmt.Println("Init")
	numero = 6
}

func main() {
	fmt.Println("Main!!!")
	fmt.Println(numero)
}
