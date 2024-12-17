package main

import (
	"log"
	"userService/internal/user/presentation"
	"userService/pkg/config"
	"userService/pkg/db"
	"userService/pkg/env"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	mongoURI := env.GetEnv("MONGO_URI", "")

	client, err := db.NewMongoDBClient(mongoURI)
	if err != nil {
		log.Fatalf("Erro ao conectar ao MongoDB: %v", err)
		return
	}
	defer client.Close()

	firebaseClient := config.NewFirebaseClient()

	presentation.RegisterUserRoutes(r, client, firebaseClient)

	// Inicia o servidor na porta 8080
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
