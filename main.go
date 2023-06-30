package main

import "fmt"

func mainx() {
	Permiso1 := 1 << 0 // Desplazamiento de 0 posiciones a la izquierda
	Permiso2 := 1 << 1 // Desplazamiento de 1 posición a la izquierda
	Permiso3 := 1 << 2 // Desplazamiento de 2 posiciones a la izquierda
	Permiso4 := 1 << 3 // Desplazamiento de 3 posiciones a la izquierda
	Permiso5 := 1 << 4 // Desplazamiento de 3 posiciones a la izquierda
	Permiso6 := 1 << 5 // Desplazamiento de 3 posiciones a la izquierda
	Permiso7 := 1 << 6 // Desplazamiento de 3 posiciones a la izquierda
	Permiso8 := 1 << 7 // Desplazamiento de 3 posiciones a la izquierda

	asignado := 7

	fmt.Println(Permiso1)
	fmt.Println(Permiso2)
	fmt.Println(Permiso3)
	fmt.Println(Permiso4)
	fmt.Println(Permiso5)
	fmt.Println(Permiso6)
	fmt.Println(Permiso7)
	fmt.Println(Permiso8)

	fmt.Println(asignado & Permiso4)
}
