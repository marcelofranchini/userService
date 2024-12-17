package config

import (
	"log"
	"userService/pkg/env"

	"github.com/go-resty/resty/v2"
)

// Estrutura para armazenar o client do Firebase, com chave da API
type FirebaseClient struct {
	Client  *resty.Client
	APIKey  string
	BaseURL string
}

func NewFirebaseClient() *FirebaseClient {
	// Obtendo a chave da API e a URL base do Firebase a partir do arquivo .env
	apiKey := env.GetEnv("FIREBASE_API_KEY", "")
	if apiKey == "" {
		log.Fatal("A chave da API do Firebase não foi configurada corretamente.")
	}

	// O Firebase não precisa de uma base URL para esse endpoint específico, então podemos omitir isso.
	client := resty.New()

	// Retorna a instância do FirebaseClient
	return &FirebaseClient{
		Client: client,
		APIKey: apiKey,
	}
}
