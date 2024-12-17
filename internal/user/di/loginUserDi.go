package di

// import (
// 	"userService/internal/user/presentation"
// 	"your_project/internal/user/app"
// 	"your_project/internal/user/domain"
// 	"your_project/internal/user/presentation"
// )

// type UserCreateDI struct {
// 	UserRepo          infra.UserRepository
// 	UserService       *app.UserService
// 	UserCreateUseCase *app.UserCreateUseCase
// 	UserHandler       *presentation.CreateUserHandler
// }

// func MakeUserCreateDI() *UserCreateDI {
// 	// Criação das dependências
// 	userRepo := infra.UserCreateRepository()                               // Repositório de usuário
// 	userService := app.NewUserService(userRepo)                      // Serviço de usuário
// 	userCreateUseCase := app.NewUserCreateUseCase(userService)       // Caso de uso de criação de usuário
// 	userHandler := presentation.CreateUserHandler(userCreateUseCase) // Handler

// 	// Retorna o contêiner de dependências preenchido
// 	return &UserCreateDI{
// 		UserRepo:          userRepo,
// 		UserService:       userService,
// 		UserCreateUseCase: userCreateUseCase,
// 		UserHandler:       userHandler,
// 	}
// }
