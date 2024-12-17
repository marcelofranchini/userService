package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(key string, defaultValue string) string {
	// Carregar as variáveis de ambiente do arquivo .env
	err := godotenv.Load()
	if err != nil {
		log.Printf("Erro ao carregar o arquivo .env: %v", err)
	}

	// Obter o valor da variável de ambiente
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
