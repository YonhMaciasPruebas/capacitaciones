package main

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
	//-----PRACTICAS-------//
	//practica_otra_forma_de_declarar_variables()
	//practica_valor_por_defecto()
	//practica_conversion_tipo_de_datos()
	//practica_paquete_fmt()
	//practica_operaciones_Aritmenticas()
	//practica_operadores_booleanos()
	//practica_condicion_if_else()
	//practica_declaracion_switch()
	//practica_bucle_for()
	//saludo := practica_uso_de_funciones_saludo("yonh")
	//fmt.Println(saludo)
	//suma_res, multiplicacion_res := practica_uso_de_funciones_calculos1(35, 42)
	//fmt.Printf("El resultado de la suma es %d y la multiplicacion es %d\n", suma_res, multiplicacion_res)
	//sum_res, multi_res := practica_uso_de_funciones_calculos2(5, 6)
	//fmt.Printf("El resultado de la suma es %d y la multiplicacion es %d\n", sum_res, multi_res)
	//practica_matrices()
	//practica_rebanadas()
	//practica_mapas()
	//practica_estructura()
	//practica_punteros_y_metodos()

	//------EJERCICIOS-----------------------//
	//ejercicio_area_perimetro_triangulo()
	//ejercicio_jugar_numero_aleatorio()
	ejercicio_gestor_de_tareas()
}
