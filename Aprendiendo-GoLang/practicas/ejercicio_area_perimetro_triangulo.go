package practicas

import (
	"fmt"
	"math"
)

func Ejercicio_area_perimetro_triangulo() {
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
