package main

import (
	"cadastro_de_clientes/config"
	"cadastro_de_clientes/models"
	"cadastro_de_clientes/routes"
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/joho/godotenv"
)
func main()  {
	config.OpenConnection()
	models.Migrate()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}
	port := os.Getenv("PORT")
	port = fmt.Sprintf(":%s", port)
	http.ListenAndServe(port,routes.ConfigureRoutes())
}
