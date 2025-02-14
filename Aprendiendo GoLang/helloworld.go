// You can edit this code!
// Click here and start typing.
package main

import "fmt"

var firstName = "yonh" //Declarar variables fuera de una fuNCION

const Pi = 3.1415 //Declarar constantes
const (
	x = 100
	y = 0b1010 //binario
	w = 0o12   //octal
	z = 0xFF   //Hexadecimal
)
const (
	Domingo = iota + 1
	Lunes
	Martes
	Miercoles
	Jueves
	Viernes
	SAbado
)

func main() {
	lastName, edad := "Macias", 30 //otra forma de declarar una variable dentro de una funcion

	fmt.Println(firstName, lastName, edad)
	fmt.Println(x, y, w, z)
	fmt.Println(Viernes)
}
