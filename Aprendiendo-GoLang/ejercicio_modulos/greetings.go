package ejercicio_modulos

import (
	"errors"
	"fmt"
	"math/rand"
)

// saluda a la perosna especifica
func Hola(name string) (string, error) {
	if name == "" {
		return "", errors.New("Nombre vacio")
	}
	//Devuelve mensaje saludo con nombre
	mensaje := fmt.Sprintf(randomSaludos(), name)
	return mensaje, nil
}

func VariosSaludos(names []string) (map[string]string, error) {
	mensajes := make(map[string]string)
	for _, name := range names {
		mensaje, error := Hola(name)
		if error != nil {
			return nil, error
		}
		mensajes[name] = mensaje
	}
	return mensajes, nil
}

// Saludos aleatorios
func randomSaludos() string {
	formats := []string{
		"¡Hola %v, bienvenido!",
		"¡Que bueno verte %v!",
		"Cordial saludo %v",
	}
	return formats[rand.Intn(len(formats))]
}
