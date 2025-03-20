package practicas

import (
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"runtime"
	"strconv"
	"time"

	"github.com/YonhMaciasPruebas/capacitaciones/Aprendiendo-Golang/ejercicio_modulos"
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

func Practica_otra_forma_de_declarar_variables() {
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

func Practica_valor_por_defecto() {
	var (
		defaultInt   int
		defaultUint  uint
		defaultFloat float32
		defaultBool  bool
		defaultSting string
	)
	fmt.Println(defaultInt, defaultUint, defaultFloat, defaultBool, defaultSting)
}

func Practica_conversion_tipo_de_datos() {

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

func Practica_paquete_fmt() {

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

func Practica_operaciones_Aritmenticas() {
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

func Practica_operadores_booleanos() {
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

func Practica_condicion_if_else() {

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

func Practica_declaracion_switch() {

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

func Practica_bucle_for() {
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
func Practica_uso_de_funciones_saludo(name string) string {
	return "Hola " + name
}

func Practica_uso_de_funciones_calculos1(a, b int) (int, int) {
	suma := a + b
	multiplicacion := a * b
	return suma, multiplicacion
}

func Practica_uso_de_funciones_calculos2(a, b int) (suma, multi int) {
	suma = a + b
	multi = a * b
	return
}

//fin Practica funciones

func Practica_matrices() {
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

func Practica_rebanadas() { //tambien se le conoce como slice
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

func Practica_mapas() {
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

func Practica_estructura() {
	type Persona struct {
		nombre string
		edad   int
		correo string
	}

	var p Persona
	p.nombre = "Yonh"
	p.edad = 28
	p.correo = "yonhmacias@outlook.com"

	p1 := Persona{"Danna", 26, "danna@gmail.com"}

	fmt.Println(p)
	fmt.Println(p1)
}

//-----TODO ESTO HACE PARTE DE LA CLASE DE PUNTEROS Y METODOS

func Practica_punteros_y_metodos() {
	var x int = 100
	var pointer *int = &x

	fmt.Println(x)
	Practica_editar_puntero(pointer)
	fmt.Println(x)

	p1 := Persona{"Danna", 26, "danna@gmail.com"} //objeto de la clase persona
	p1.saludar()
}

func Practica_editar_puntero(a *int) {
	*a = 20
}

type Persona struct { //ESTRUCTURA TIPO CLASE CON ATRIBUTOS
	nombre string
	edad   int
	correo string
}

func (c *Persona) saludar() { // METODO DE LA CLASE SE USA PUNTEROS
	fmt.Printf("Hola, mi nombre es %s y tengo %d años\n", c.nombre, c.edad)
}

// -----HASTA AQUI HACE PARTE DE LA CLASE DE PUNTEROS Y METODOS

//------------------------------------------------------------------------------
//SECCION 6 CONTROL DE ERRORES

// comienza practica manejo de errores
func Practica_manejo_de_errores() {
	/*str := "a"
	num, error := strconv.Atoi(str)
	if error != nil {
		fmt.Println("Error: ", error)
		return
	}
	fmt.Println("Número: ", num)*/

	resultado, error := calcular_division(10, 0)
	if error != nil {
		fmt.Println("Error: ", error)
		return
	}
	fmt.Println("Resultado: ", resultado)
}
func calcular_division(dividendo, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("No se puede dividir por cero")
		//return 0, fmt.Errorf("No se puede dividir por cero") // asi tambien se puede

	}
	return dividendo / divisor, nil
}

//finaliza practica manejo de errores

func Practica_uso_defer() {
	file, err := os.Create("soy_un_archivo.txt")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer file.Close() //esto hace que sea lo ultimo que se ejecute antes de que finalice main

	_, err = file.Write([]byte("Hola, mundo soy yonh"))
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
}

// inicia practica uso de panic y recover
func Practica_uso_panic_y_recover() {
	divide_panic(67, 4)
	divide_panic(50, 5)
	divide_panic(3, 0)
	divide_panic(100, 6)

}

func divide_panic(dividendo, divisor int) {
	//recupera el panic para permitir que la ejecucion no finalice sino trate el dato dañado y continue con los demas registros
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	validacionCero(divisor)
	fmt.Println(dividendo / divisor)
}

func validacionCero(divisor int) {
	if divisor == 0 {
		panic("No es posible divir por cero")
	}
}

//finaliza practica uso de panic y recover

func Practica_registro_de_errores() {
	log.SetPrefix("Practica_registro_de_errores(): ")                                 //colocar un prefijo, suele ser el nombre de la funcion para que en la traza quede como parte del log
	file, error := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644) //nombrearchivo,si no existe lo crea, escribe registros si ya tiene,solo lo apertura en escritura,permisos de ejecucion
	if error != nil {
		log.Fatal(error)
		//log.Fatal("OJO son un mensaje fatal")                                             //con este finaliza la ejecucion del programa
		//log.Panic("OJO son un mensaje de panic")                                          //con este finaliza pero deja una traza de donde ocurrieron los errores

	}
	defer file.Close()

	log.SetOutput(file) //para que no escriba en consola sino en el archivo info.log
	log.Print("¡soy un registro del log!")
	log.Print("Este es un mensaje de registro")
	log.Println("Este es otro mensaje de registro")

}

func Practica_modulos() {
	log.SetPrefix("ejercicio_modulos: ")
	//log.SetFlags(0) // elimina todos los datos adicionales del log como la fecha y la hora

	mensaje, error := ejercicio_modulos.Hola("Yonh")
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println(mensaje)

	nombres := []string{"Yonh", "Danna", "Juan", "Maria", "Tatiana", "Hernan"}
	map_mensajes, error := ejercicio_modulos.VariosSaludos(nombres)
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println(map_mensajes)
}
