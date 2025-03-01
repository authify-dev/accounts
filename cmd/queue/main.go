package main

import (
	"accounts/internal/common/controllers/queue"
	"fmt"
	"log"
	"time"
	// Asegúrate de que el paquete "queue" esté en el path correcto.
)

func main() {
	// Crear una instancia del controlador de colas
	qc := queue.NewQueueController()

	// Consumir mensajes de la cola "new-users"
	msgs, err := qc.ConsumeFromQueue("users_registered", "send_code_of_verification")
	if err != nil {
		log.Fatalf("Error consumiendo la cola: %v", err)
	}

	// Procesar los mensajes recibidos en una gorutina
	go func() {
		for msg := range msgs {
			time.Sleep(5 * time.Second)
			fmt.Printf("Mensaje recibido: %s\n", msg.Body)
		}
	}()

	// Mantener el programa en ejecución para poder consumir mensajes.
	fmt.Println("Esperando mensajes. Para salir presiona CTRL+C")
	for {
		time.Sleep(1 * time.Second)
	}
}
