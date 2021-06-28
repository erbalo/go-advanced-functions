package main

import (
	"fmt"
	"math"
)

// la implementacion de la interfaz es implicita
type forma interface {
	obtenerNombre() string
	area() float64
}

type rectangulo struct {
	name   string
	altura float64
	ancho  float64
}

type circulo struct {
	name  string
	radio float64
}

func (r rectangulo) area() float64 {
	return r.altura * r.ancho
}

func (r rectangulo) obtenerNombre() string {
	return r.name
}

func (c circulo) area() float64 {
	return math.Pi * math.Pow(c.radio, 2)
}

func (c circulo) obtenerNombre() string {
	return c.name
}

func escribirArea(f forma) {
	fmt.Printf("El área de la forma [%s] es %0.2f\n", f.obtenerNombre(), f.area())
}

func main() {
	r := rectangulo{"rectangulo", 10, 15}
	escribirArea(r)

	c := circulo{"circulo", 5.83}
	escribirArea(c)

}
