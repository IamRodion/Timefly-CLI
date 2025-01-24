package utils

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

// Funciones -----------------------------------------------------------------------------------------------------

// Función que limpia la terminal (ejecuta un "clear" en la terminal)
func CleanCli() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// Función que realiza un salto de línea
func JumpLine() {
	fmt.Println("")
}

/*
Toma varios argumentos y ejecuta un comando con ellos de la siguiente forma:

	dir string      => // Un string con la ruta donde se ejecutará el comando
	command string  => // Un string con el comando a ejecutar
	arg ...string   => // Una cantidad indefinida de strings para utilizar como argumentos del comando a ejecutar
*/
func RunCommand(dir string, command string, arg ...string) {
	cmd := exec.Command(command, arg...)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr // Esto es útil para ver errores del comando
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error al crear el proyecto:", err)
	}

}

// Structs -----------------------------------------------------------------------------------------------------

/*
Struct para mostrar mensajes mientras hay tareas en ejecución

	Message string 	// El mensaje que se mostrará antes del símbolo que indica que se está ejecutando una tarea
	Index   int 	// El index del primer símbolo a mostrar
*/
type Loading struct {
	Message string
	Index   int
}

// Método del struct Loading que muestra un mensaje indicando que se está realizando una tarea
func (loading *Loading) ShowMessage() {
	symbols := [12]string{"⠁", "⠉", "⠙", "⠹", "⠸", "⠰", "⠠", "⠤", "⠦", "⠧", "⠇", "⠃"}
	fmt.Print(loading.Message, symbols[loading.Index%len(symbols)], "\r")
	loading.Index = loading.Index % len(symbols)
	loading.Index += 1
	time.Sleep(70 * time.Millisecond)
}
