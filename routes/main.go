package routes

import (
	"cadastro_de_clientes/handlers"
	"net/http"

	"github.com/go-chi/chi"
)

func ConfigureRoutes() *chi.Mux{
	router := chi.NewRouter()
	router.Route("/api", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("API Principal"))
		})
		r.Mount("/user", UsersRoutes())
		r.Mount("/session", SessionRoutes())
		r.Mount("/client", ClientsRoutes())
	})
	return router
}

func UsersRoutes() chi.Router{
	router := chi.NewRouter()
	router.Use(handlers.AuthenticateToken)
	router.Get("/", handlers.ProfileUser)
	return router
}

func SessionRoutes() chi.Router{
	router := chi.NewRouter()
	router.Post("/login", handlers.LoginUser)
	router.Post("/register", handlers.RegisterUser)
	return router
}
func ClientsRoutes() chi.Router{
	router := chi.NewRouter()
	router.Use(handlers.AuthenticateToken)
	router.Get("/", handlers.Get)
	router.Post("/", handlers.Post)
	return router
}
