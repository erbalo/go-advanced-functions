package main

import "fmt"

func log(inf interface{}) {
	fmt.Println(inf)
}

func alumnoEsAprobado(cal1, cal2 float32) bool {
	defer fmt.Println("Imprimiendo media...") // skip
	defer fmt.Println("Otra cosa")
	defer log("Epale!")
	fmt.Println("Entrando a funcion para verificar si un alumno es aprobado")

	media := (cal1 + cal2) / 2

	if media >= 6 {
		return true
	}

	return false
}

func main() {
	fmt.Println("Funciones defer")

	fmt.Println(alumnoEsAprobado(5, 5))
}
