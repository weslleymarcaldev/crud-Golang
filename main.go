package main

import (
	"log"
	"net/http"
	"projetoGo/config"
	"projetoGo/handlers"
)

func main() {
	// Configuração RabbitMQ
	rabbitConn, err := config.ConnectRabbitMQ()
	if err != nil {
		log.Fatalf("Falha ao conectar ao RabbitMQ: %v", err)
	}
	defer rabbitConn.Close()

	rabbitChannel, err := rabbitConn.Channel()
	if err != nil {
		log.Fatalf("Falha ao abrir um canal RabbitMQ: %v", err)
	}
	defer rabbitChannel.Close()

	// Declara a fila no RabbitMQ
	queue, err := config.DeclareQueue(rabbitChannel, "user_queue")
	if err != nil {
		log.Fatalf("Falha ao declararRabbitMQ queue: %v", err)
	}

	// Configuração Memcache
	memcacheClient := config.ConnectMemcache()

	// Inicializando Handlers
	handlers.InitHandlers(memcacheClient, rabbitChannel, queue.Name)

	// Rotas
	http.HandleFunc("/create", handlers.CreateUser)
	http.HandleFunc("/get", handlers.GetUser)

	// Servidor
	log.Println("Servidor em execução na porta 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
