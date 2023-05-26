package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenConnection() (*gorm.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}
	connection	 := os.Getenv("CONNECTION")
	conn, err := gorm.Open(postgres.Open(connection))
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	log.Println("Conexão com o banco de dados estabelecida com sucesso")
	return conn, nil
}
