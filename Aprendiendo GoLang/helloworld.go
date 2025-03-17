// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"math"
	"strconv"
	"time"
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
	//practica_otra_forma_de_declarar_variables()
	//practica_valor_por_defecto()
	//practica_conversion_tipo_de_datos()
	//practica_paquete_fmt()
	//practica_operaciones_Aritmenticas()
	//ejercicio_area_perimetro_triangulo()
	//practica_operadores_booleanos()
	practica_condicion_if_else()
}

func practica_otra_forma_de_declarar_variables() {
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

func practica_valor_por_defecto() {
	var (
		defaultInt   int
		defaultUint  uint
		defaultFloat float32
		defaultBool  bool
		defaultSting string
	)
	fmt.Println(defaultInt, defaultUint, defaultFloat, defaultBool, defaultSting)
}

func practica_conversion_tipo_de_datos() {

	var integer16 int16 = 50
	var integer32 int32 = 100

	fmt.Println(int32(integer16) + integer32)

	s := "200"
	num, _ := strconv.Atoi(s)

	fmt.Println(num + num)

	n := 42
	n_string := strconv.Itoa(n)
	fmt.Println(n_string + n_string)
}

func practica_paquete_fmt() {

	var nombre string
	var edad int

	fmt.Print("Ingrese su nombre: ")
	fmt.Scanln(&nombre)
	fmt.Print("Ingrese su edad: ")
	fmt.Scanln(&edad)

	//fmt.Printf("Hola, me llamo %s y tengo %d años.\n", nombre, edad)
	saludo := fmt.Sprintf("Hola, me llamo %s y tengo %d años.\n", nombre, edad)
	fmt.Println(saludo)

	fmt.Printf("El tipo de la variable nombre es: %T\n", nombre)
	fmt.Printf("El tipo de la variable edad es: %T\n", edad)
}

func practica_operaciones_Aritmenticas() {
	a := 10.0
	b := 3.0

	b++
	b--

	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(int(a) % int(b))

	c := 5
	c += 2
	fmt.Println(c)
	fmt.Println(math.E)
	fmt.Println(math.Sqrt(25))

}

func ejercicio_area_perimetro_triangulo() {
	//declaracion de variables
	var altura, lado, hipotenusa, area, perimetro float32
	const precision_decimal = 2

	//Entrada de datos
	fmt.Print("Ingrese el lado del triangulo: ")
	fmt.Scanln(&lado)
	fmt.Print("Ingrese la altura del triangulo: ")
	fmt.Scanln(&altura)

	//Calculos

	hipotenusa = float32(math.Sqrt(math.Pow(float64(lado), 2) + math.Pow(float64(altura), 2)))
	area = (lado * altura) / 2
	perimetro = hipotenusa + lado + altura

	//Salida de datos
	fmt.Printf("Area: %.*f\n", precision_decimal, area)
	fmt.Printf("Perimetro: %.*f\n", precision_decimal, perimetro)
}

func practica_operadores_booleanos() {
	// Comparación de números
	fmt.Println(1 == 1) // true
	fmt.Println(1 != 2) // true
	fmt.Println(2 < 3)  // true
	fmt.Println(3 > 4)  // false
	fmt.Println(4 <= 4) // true
	fmt.Println(5 >= 6) // false

	// Comparación de cadenas
	fmt.Println("hola" == "hola")  // true
	fmt.Println("hola" != "adios") // true
	fmt.Println("abc" < "def")     // true
	fmt.Println("ghi" > "fgh")     // true
	fmt.Println("hij" <= "hij")    // true
	fmt.Println("klm" >= "klmno")  // false

	// Comparación de booleanos
	fmt.Println(true == true)           // true
	fmt.Println(false != true)          // true
	fmt.Println(true && false == false) // true
	fmt.Println(true || false == true)  // true

	//operador and
	x := true
	y := false
	z := x && y
	fmt.Println(z) // Imprime false

	//operador or
	x1 := true
	y1 := false
	z1 := x1 || y1
	fmt.Println(z1) // Imprime true

	//operador not logico
	x3 := true
	z3 := !x3
	fmt.Println(z3) // Imprime false
}

func practica_condicion_if_else() {

	tiempo := time.Now()
	hora := tiempo.Hour()

	if hora < 12 {
		fmt.Println("¡Mañana!")
	} else if hora < 18 {
		fmt.Println("¡Tarde!")
	} else {
		fmt.Println("¡Noche!")
	}
	//////////////Es lo mismo //////////////
	if t := time.Now(); t.Hour() < 12 {
		fmt.Println("¡Mañana!")
	} else if t.Hour() < 18 {
		fmt.Println("¡Tarde!")
	} else {
		fmt.Println("¡Noche!")
	}
}
