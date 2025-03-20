package practicas

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

// Estrcutura de contactos
type Contactos struct {
	Nombre   string `json:"nombre"`
	Correo   string `json:"correo"`
	Telefono string `json:"telefono"`
}

// Guardar contactos en un archivo json
func guardarContactosEnArchivo(contactos []Contactos) error {
	file, err := os.Create("contactos.json")
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file) //instancia para codificar el dato tipo slice a json
	err = encoder.Encode(contactos)  // codifica el dato que esta en tipo slice a json
	if err != nil {
		return err
	}
	return nil
}

// Cargar contactos de un archivo json
func cargaContactosDeUnArchivo(contactos *[]Contactos) error {
	file, err := os.Open("contactos.json")
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file) //instancia para decodificar el dato tipo json a slice
	err = decoder.Decode(&contactos) // decodifica el dato que esta en tipo json a slice y lo guarda en la memoria de la estrcutura

	if err != nil {
		return err
	}
	return nil

}

func Ejercicio_gestor_de_contactos() {
	var contactos []Contactos
	err := cargaContactosDeUnArchivo(&contactos)
	if err != nil {
		fmt.Println("no se logro cargar contactos: ", err)
	}

	//Crear una instancia del paquete bufio para leer por teclado
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("---- GESTOR DE CONTACTOS ----\n",
			"1. Agregar contacto\n",
			"2. Mostrar todos los contactos\n",
			"3. Salir\n",
			"Eliga una opción: ")

		//Leer la opcion del usuario
		var opcion int
		_, err = fmt.Scanln(&opcion)
		if err != nil {
			fmt.Println("Error al leer opcion del usuario: ", err)
			return
		}
		//Manejar opcion del usuario
		switch opcion {
		case 1:
			var c Contactos
			fmt.Print("Ingrese el nombre de contacto: ")
			c.Nombre, _ = reader.ReadString('\n')
			fmt.Print("Ingrese el correo de contacto: ")
			c.Correo, _ = reader.ReadString('\n')
			fmt.Print("Ingrese número de contacto: ")
			c.Telefono, _ = reader.ReadString('\n')

			//agregar contactos al slice
			contactos = append(contactos, c)
			if err := guardarContactosEnArchivo(contactos); err != nil {
				fmt.Println("Error al guardar contactos: ", err)
			} else {
				fmt.Println("Contacto agregado correctamente")
			}
		case 2:
			fmt.Println("|---------------------------------------------------|")
			for index, contacto := range contactos {
				fmt.Printf("%d. Nombre: %s Email: %s Telefono: %s\n", index+1, contacto.Nombre, contacto.Correo, contacto.Telefono)
			}
			fmt.Println("|---------------------------------------------------|")
		case 3:
			fmt.Println("Finalizando programa...")
			return
		default:
			fmt.Println("Opción inválida")
		}
	}
}
