// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"math"
)

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

var numero uint8 = 2 //numeros enteros positivos maximo 8 bits
var numero2 int8 = 2 //numeros enteros negativos y positivos maximo 8 bits
var numerofloat float32 = 4.5
var datobool bool = true

func main() {
	lastName, edad := "Macias", 30 //otra forma de declarar una variable dentro de una funcion

	fmt.Println(firstName, lastName, edad)
	fmt.Println(x, y, w, z)
	fmt.Println(Viernes)

	fmt.Println(math.MaxInt64)

	fullname := "Yonh \t(alias \"mathews\")\n"
	fmt.Println(fullname)

	var a byte = 'a'
	fmt.Println(a)

	s := "hola"
	fmt.Println(s[0])

	var corazon rune = '$'
	fmt.Println(corazon)
}
