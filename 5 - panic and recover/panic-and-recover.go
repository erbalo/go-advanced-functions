package main

import "fmt"

func recuperarEjecucion() {
	if r := recover(); r != nil {
		fmt.Println("Intentando recuperar la ejecución")
		fmt.Println(r)
	}
}

func alumnoEsAprobado(cal1, cal2 float32) bool {
	defer recuperarEjecucion()
	fmt.Println("Entrando a funcion para verificar si un alumno es aprobado")

	media := (cal1 + cal2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("La media es exactamente 6!") // -->
}

func main() {
	fmt.Println(alumnoEsAprobado(6, 6))
	fmt.Println("Post-ejecucion")
}
