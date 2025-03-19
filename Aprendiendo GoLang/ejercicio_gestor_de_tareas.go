package main

import (
	"bufio"
	"fmt"
	"os"
)

// estructura de una tarea
type Tarea struct {
	nombre      string
	descripcion string
	completado  bool
}

// estructura de una lista de tareas que va a contener las tareas
type ListaTareas struct {
	tareas []Tarea
}

// metodo de agregar tareas a la lista
func (lista_tareas *ListaTareas) agregarTarea(t Tarea) {
	lista_tareas.tareas = append(lista_tareas.tareas, t)

}

// Metodo para marcar completada una tarea
func (lista_tareas *ListaTareas) marcarCompletado(posicion int) {
	lista_tareas.tareas[posicion].completado = true
}

// Metodo para editar una tarea
func (lista_tareas *ListaTareas) editarTareas(posicion int, t Tarea) {
	lista_tareas.tareas[posicion] = t
}

// Metodo para eliminar una tarea
func (lista_tareas *ListaTareas) eliminarTareas(posicion int) {
	lista_tareas.tareas = append(lista_tareas.tareas[:posicion], lista_tareas.tareas[posicion+1:]...)
}

// Aqui creo una interfaz
func ejercicio_gestor_de_tareas() {
	//Instancia de lista de tareas
	lista := ListaTareas{}
	menu_principal(lista)

}

func menu_principal(lista ListaTareas) {
	leer := bufio.NewReader(os.Stdin)
	for {
		var opcion int
		fmt.Println("Seleccione una opción:\n",
			"1. Agregar tarea\n",
			"2. Marcar tarea como completada\n",
			"3. Editar tarea\n",
			"4. Eliminar tarea\n",
			"5. Salir")

		fmt.Print("Ingrese una opción: ")
		fmt.Scanln(&opcion)
		switch opcion {
		case 1:
			var t Tarea
			fmt.Print("Ingrese el nombre de la tarea: ")
			t.nombre, _ = leer.ReadString('\n')
			fmt.Print("Ingrese la descripción de la tarea: ")
			t.descripcion, _ = leer.ReadString('\n')
			lista.agregarTarea(t)
			fmt.Println("Tarea agregada correctamente")
		case 2:
			var posicion_tarea int
			fmt.Print("Ingrese el número de la tarea que desea marcar como completada: ")
			fmt.Scanln(&posicion_tarea)
			lista.marcarCompletado(posicion_tarea)
			fmt.Println("Tarea marcada como completada correctamente")
		case 3:
			var posicion_tarea int
			var t Tarea
			fmt.Print("Ingrese el número de la tarea que desea actualizar: ")
			fmt.Scanln(&posicion_tarea)
			fmt.Print("Ingrese el nombre de la tarea: ")
			t.nombre, _ = leer.ReadString('\n')
			fmt.Print("Ingrese la descripción de la tarea: ")
			t.descripcion, _ = leer.ReadString('\n')
			lista.editarTareas(posicion_tarea, t)
			fmt.Println("Tarea actualizada correctamente")
		case 4:
			var posicion_tarea int
			fmt.Print("Ingrese el número de la tarea que desea eliminar: ")
			fmt.Scanln(&posicion_tarea)
			lista.eliminarTareas(posicion_tarea)
			fmt.Println("Tarea eliminada correctamente")
		case 5:
			fmt.Println("Finalizando programa...")
			return
		default:
			fmt.Println("Opción incorrecta")
		}
		mostrarListaTareas(lista)
	}
}

func mostrarListaTareas(lista ListaTareas) {
	// Se enlistan todas las tareas
	fmt.Println("Lista de tareas")
	fmt.Println("|------------------------------------------------------------------|")
	for index, tarea := range lista.tareas {
		fmt.Printf("%d. %s - %s - Completado: %t\n", index, tarea.nombre, tarea.descripcion, tarea.completado)
	}
	fmt.Println("|------------------------------------------------------------------|")
}
