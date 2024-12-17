package infra

import (
	"fmt"
	"log"
	"userService/internal/user/domain"
	"userService/pkg/config"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
)

type FirebaseRepository struct {
	FirebaseClient *config.FirebaseClient
	FirebaseApp    *firebase.App
	FirebaseAuth   *auth.Client
}

func NewFirebaseRepository(firebaseClient *config.FirebaseClient) (*FirebaseRepository, error) {
	opt := option.WithCredentialsFile("/Users/marcelofranchini/Desktop/AppGo/userService/internal/user/infra/firebaseconfig.json") // Caminho para o arquivo JSON

	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, fmt.Errorf("erro ao inicializar o Firebase: %v", err)
	}

	authClient, err := app.Auth(context.Background())
	if err != nil {
		return nil, fmt.Errorf("erro ao obter o cliente de autenticação do Firebase: %v", err)
	}

	return &FirebaseRepository{
		FirebaseClient: firebaseClient,
		FirebaseApp:    app,
		FirebaseAuth:   authClient,
	}, nil
}

func (r *FirebaseRepository) CreateUser(user *domain.UserCreateRequest) error {
	params := (&auth.UserToCreate{}).
		Email(user.Email).
		Password(user.Password).
		EmailVerified(false)

	newUser, err := r.FirebaseAuth.CreateUser(context.Background(), params)
	if err != nil {
		log.Printf("Erro ao criar usuário no Firebase: %v", err)
		return err
	}

	log.Printf("Usuário criado com sucesso: %v", newUser.UID)
	return nil
}
