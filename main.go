package main

import (
	"cadastro_de_clientes/config"
	"cadastro_de_clientes/handlers"
	"cadastro_de_clientes/models"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)
func main()  {
	config.OpenConnection()
	models.Migrate()
	r := chi.NewRouter()
	r.Post("/user/register", handlers.RegisterUser)
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}
	port := os.Getenv("PORT")
	port = fmt.Sprintf(":%s", port)
	http.ListenAndServe(port,r)
}
