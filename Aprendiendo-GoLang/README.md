# Saludos en GO

Este paquete proporciona una forma simple de obtener saludos personalizados en Go.

## Instalación

Ejecuta el siguiente comando para instalar el paquete:
 ```bash
 go get -u github.com/YonhMaciasPruebas/capacitaciones/Aprendiendo-Golang
 ```

## Uso
Aquí se muestra un ejemplo de cómo se utiliza el paquete

```go
package main

import (
    "fmt"
	"github.com/YonhMaciasPruebas/capacitaciones/Aprendiendo-Golang/practicas"
)

func main() {

	mensaje, error := ejercicio_modulos.Hola("Yonh")
	if error != nil {
		fmt.Println("Ocurrió un error: ", error)
	}
	fmt.Println(mensaje)

	nombres := []string{"Yonh", "Danna", "Juan", "Maria", "Tatiana", "Hernan"}
	map_mensajes, error := ejercicio_modulos.VariosSaludos(nombres)
	if error != nil {
		fmt.Println("Ocurrió un error: ", error)
	}
	fmt.Println(map_mensajes)
}