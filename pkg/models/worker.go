package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// Worker representa la estructura de cada worker en la respuesta de la API
type Worker struct {
	ID         int    `json:"id"`
	Firstname  string `json:"firstname"`
	Lastname   string `json:"lastname"`
	Email      string `json:"email"`
	EmployeeID string `json:"employe_id"`
	Department string `json:"department"`
	HireDate   string `json:"hire_date"`
	Active     bool   `json:"active"`
}

// Método que muestra los datos del struct worker
func (worker Worker) ShowData() {
	fmt.Printf("ID: %s, Nombre: %s %s, Email: %s, Departamento: %s\n", worker.EmployeeID, worker.Firstname, worker.Lastname, worker.Email, worker.Department)
}

func (worker Worker) RegisterEntry(entry_type string) {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error al cargar el archivo .env: %v", err)
	}

	// Obtener las variables de entorno
	apiURL := os.Getenv("API_URL")
	url := fmt.Sprintf("%sTimeEntry/", apiURL)

	timeEntry := TimeEntry{EntryType: entry_type, Worker: worker.ID}

	// Serializar el struct TimeEntry a JSON
	jsonData, err := json.Marshal(timeEntry)
	if err != nil {
		fmt.Println("Error al serializar el JSON:", err)
		return
	}

	// Realizar la solicitud POST
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error al realizar la solicitud POST:", err)
		return
	}
	defer resp.Body.Close()

	// Verificar el código de respuesta
	if resp.StatusCode != http.StatusCreated {
		fmt.Printf("Error: código de respuesta %d\n", resp.StatusCode)
		return
	}

	switch {
	case timeEntry.EntryType == "IN":
		fmt.Printf("Entrada de %s %s registrada exitosamente\n", worker.Firstname, worker.Lastname)
	case timeEntry.EntryType == "OUT":
		fmt.Printf("Salida de %s %s registrada exitosamente\n", worker.Firstname, worker.Lastname)
	}
	fmt.Scanln()
}

func GetWorkers(ch chan []Worker) {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error al cargar el archivo .env: %v", err)
	}

	// Obtener las variables de entorno
	apiURL := os.Getenv("API_URL")
	url := fmt.Sprintf("%sWorker/", apiURL)

	// Realizar la petición GET
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error al realizar la petición:", err)
		return
	}
	defer resp.Body.Close()

	// Leer el cuerpo de la respuesta
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error al leer la respuesta:", err)
		return
	}

	// Convertir el cuerpo de la respuesta en un slice de estructuras Worker
	var workers []Worker
	if err := json.Unmarshal(body, &workers); err != nil {
		fmt.Println("Error al deserializar el JSON:", err)
		return
	}

	ch <- workers
	close(ch)
}
