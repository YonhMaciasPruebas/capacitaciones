package main

import (
	"fmt"
	"math/rand"
)

func ejercicio_jugar_numero_aleatorio() {
	numero_aleatorio := rand.Intn(10) + 1
	var numero_ingresado, intentos int
	const maxintentos = 3
	for intentos < maxintentos {
		intentos++
		fmt.Printf("Ingresa un número (intentos restantes %d): ", maxintentos-intentos+1)
		fmt.Scanln(&numero_ingresado)
		switch true {
		case numero_ingresado == numero_aleatorio:
			fmt.Println("¡FELICITACIONES!,Encontraste el número")
			jugar_nuevamente()
			return
		case numero_ingresado < numero_aleatorio:
			fmt.Println("El número a adivinar es mayor")
		case numero_ingresado > numero_aleatorio:
			fmt.Println("El número a adivinar es menor")
		}
	}
	fmt.Println("¡LO SIENTO!, se acabaron los intentos, el número era: ", numero_aleatorio)
	jugar_nuevamente()
}

func jugar_nuevamente() {
	var decision string
	fmt.Println("¿Desea jugar nuevamente? (si/no)")
	fmt.Scanln(&decision)

	switch decision {
	case "si":
		ejercicio_jugar_numero_aleatorio()
	case "no":
		fmt.Println("Gracias por jugar.")
	default:
		fmt.Println("Decisión inválida, intente de nuevo.")
		jugar_nuevamente()
	}
}
