package presentation

import (
	"log"
	"userService/internal/user/di"
	"userService/pkg/config"
	"userService/pkg/db"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine, client *db.MongoDBClient, firebaseClient *config.FirebaseClient) {
	userCreateDI, err := di.MakeUserCreateDI(client, firebaseClient)
	if err != nil {
		log.Fatalf("Erro ao criar DI para criação de usuário: %v", err)
		return
	}

	userRoutes := r.Group("/user")
	{
		userRoutes.POST("/create", userCreateDI.UserHandler.CreateUser)
		// userRoutes.POST("/login", userHandler.Login)
	}
}
