package routes

import (
	"cadastro_de_clientes/handlers"
	// "cadastro_de_clientes/routes"
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
		r.Mount("/email", EmailsRoutes())
		r.Mount("/cellphone", CellPhoneRoutes())
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
func EmailsRoutes() chi.Router{
	router := chi.NewRouter()
	router.Use(handlers.AuthenticateToken)
	router.Post("/{id}", handlers.PostEmail)
	router.Get("/",handlers.GetEmails)
	router.Delete("/{id}", handlers.DeleteEmail)
	return router
}
func CellPhoneRoutes() chi.Router{
	router := chi.NewRouter()
	router.Use(handlers.AuthenticateToken)
	router.Post("/{id}", handlers.PostCellphone)
	router.Get("/",handlers.GetCellphones)
	router.Delete("/{id}", handlers.DeleteCellphone)
	return router
}
