package main

import (
	"fmt"
	"math"
	"runtime"
	"strconv"
	"time"
)

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

func practica_declaracion_switch() {

	os := runtime.GOOS //saber el sistema operativo
	switch os {
	case "windows":
		fmt.Println("Go run -> Windows")
	case "linux":
		fmt.Println("Go run -> Linux")
	case "darwin":
		fmt.Println("Go run -> macOS")
	default:
		fmt.Println("Go run -> Otro OS")
	}

	//es lo mismo//

	switch os := runtime.GOOS; os {
	case "windows":
		fmt.Println("Go run -> Windows")
	case "linux":
		fmt.Println("Go run -> Linux")
	case "darwin":
		fmt.Println("Go run -> macOS")
	default:
		fmt.Println("Go run -> Otro OS")
	}

	//similar al caso time del if//

	switch t := time.Now(); {
	case t.Hour() < 12:
		fmt.Println("¡Mañana!")
	case t.Hour() < 18:
		fmt.Println("¡Tarde!")
	default:
		fmt.Println("¡Noche!")
	}
}

func practica_bucle_for() {
	var i int
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	for j := 1; j <= 10; j++ {
		if j == 6 {
			//break
			continue
		}
		fmt.Println(j)
	}
}

// Practica de funciones
func practica_uso_de_funciones_saludo(name string) string {
	return "Hola " + name
}

func practica_uso_de_funciones_calculos1(a, b int) (int, int) {
	suma := a + b
	multiplicacion := a * b
	return suma, multiplicacion
}

func practica_uso_de_funciones_calculos2(a, b int) (suma, multi int) {
	suma = a + b
	multi = a * b
	return
}

//fin Practica funciones

func practica_matrices() {
	var matriz [5]int
	matriz[3] = 5

	fmt.Println(matriz)

	var a = [5]int{1, 2, 3, 4, 5}
	fmt.Println(a)

	var b = [...]string{"a", "b", "c", "d", "f", "f", "g", "h"}
	for i := 0; i < len(b); i++ {
		fmt.Println(b[i])
	}

	for index, valor := range a {
		fmt.Printf("Indice: %d, Valor: %d\n", index, valor)
	}

	var c = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(c)

}

func practica_rebanadas() {
	diasSemana := []string{"Domingo", "Lunes", "Martes", "Miercoles", "Jueves", "Viernes", "Sabado"}
	fmt.Println(diasSemana)
	diasRebanada := diasSemana[2:4]
	//agregar elementos
	diasRebanada = append(diasRebanada, "YONH", "danna", "Juan")
	fmt.Println(diasRebanada)

	//eliminar elementos
	diasRebanada = append(diasRebanada[:2], diasRebanada[3:]...)
	fmt.Println(diasRebanada)

	fmt.Println(len(diasRebanada))
	fmt.Println(cap(diasRebanada))

	//crear rebanadas
	nombres := make([]string, 5, 10)
	nombres[1] = "Juan"
	fmt.Println(nombres)

	//copiar rebanadas
	rebanada1 := []int{2, 4, 5, 6, 7}
	rebanada2 := make([]int, 5)
	copy(rebanada2, rebanada1)
	fmt.Println("rebanda1: ", rebanada1)
	fmt.Println("rebanda2: ", rebanada2)
}

func practica_mapas() {
	colores := map[string]string{
		"rojo":  "#FF0000",
		"verde": "#00FF00",
		"azul":  "#0000FF",
	}
	colores["negro"] = "#000000" //agregar un elemento
	fmt.Println(colores)
	fmt.Println(colores["azul"])

	valor, ok := colores["verde"]
	fmt.Println(valor, ok)
	valor, ok = colores["blanco"]
	fmt.Println(valor, ok)

	if valor, ok := colores["Blanco"]; ok {
		fmt.Println(valor)
	} else {
		fmt.Println("No existe la clave")
	}

	delete(colores, "azul") //borrar un elemento
	fmt.Println(colores)

	for clave, dato := range colores {
		fmt.Printf("Clave: %s, Valor: %s\n", clave, dato)
	}
}
