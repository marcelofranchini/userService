package di

import (
	"fmt"
	"userService/internal/user/app/useCase"
	"userService/internal/user/domain"
	"userService/internal/user/infra"
	"userService/internal/user/presentation/handler"
	"userService/pkg/config"
	"userService/pkg/db"
)

type UserCreateDI struct {
	UserRepo          domain.UserCreateRepositoryInterface
	UserFirebaseRepo  domain.FirebaseRepositoryInterface
	UserCreateUseCase *useCase.CreateUserUseCase
	UserHandler       *handler.CreateUserHandler
}

func MakeUserCreateDI(client *db.MongoDBClient, firebaseClient *config.FirebaseClient) (*UserCreateDI, error) {
	database := client.Client.Database("users")
	userRepo := infra.UserCreateRepository(database)

	useFirebaseRepo, err := infra.NewFirebaseRepository(firebaseClient)
	if err != nil {
		return nil, fmt.Errorf("erro ao criar o repositório Firebase: %v", err)
	}

	userCreateUseCase := useCase.NewCreateUser(userRepo, useFirebaseRepo)
	userHandler := handler.NewCreateUserHandler(userCreateUseCase)

	return &UserCreateDI{
		UserRepo:          userRepo,
		UserCreateUseCase: userCreateUseCase,
		UserHandler:       userHandler,
	}, nil
}
