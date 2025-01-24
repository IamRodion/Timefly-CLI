package main

import (
	"fmt"

	"github.com/IamRodion/Timefly-CLI/pkg/models"
	"github.com/IamRodion/Timefly-CLI/pkg/utils"
)

func GetTypeEntry() string {
	for {
		utils.CleanCli()
		fmt.Println("\t\tTimefly CLI")
		fmt.Print("\n[1] Registrar Entrada\n[2] Registrar Salida\n")

		var typeEntry string
		fmt.Print("\n[?] Indique una opción: ")
		fmt.Scanln(&typeEntry)

		switch {
		case typeEntry == "1":
			return "IN"
		case typeEntry == "2":
			return "OUT"
		default:
			fmt.Printf("[!] La opción '%s' no es correcta", typeEntry)
			fmt.Scanln()
		}
	}
}

func GetWorkerEmployeeID(typeEntry string) {
	for {
		loading := utils.Loading{Message: "Cargando ", Index: 0}

		ch := make(chan []models.Worker)
		go models.GetWorkers(ch)

		utils.CleanCli()
		fmt.Println("\t\tTimefly CLI")

		var employeeID string
		fmt.Print("\n[?] Indique su número de empleado: ")
		fmt.Scanln(&employeeID)

		select {
		case workers := <-ch:
			for _, worker := range workers {
				if worker.EmployeeID == employeeID {
					worker.RegisterEntry(typeEntry)
					return
				}
			}
			fmt.Printf("[!] No se encontró el número de empleado '%s'", employeeID)
			fmt.Scanln()

		default:
			loading.ShowMessage()
		}
	}
}

func main() {
	for {
		GetWorkerEmployeeID(GetTypeEntry())
	}
}
