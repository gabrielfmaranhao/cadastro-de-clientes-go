package handlers

import (
	"cadastro_de_clientes/models"
	"encoding/json"
	"fmt"
	"net/http"
)
type register struct {
	Username string
	Cpf string
	Password string
}
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user register
	w.Header().Set("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&user)
	newUser,err := models.NewUser(user.Username, user.Cpf, user.Password)
	fmt.Println(err)
	if err != nil {
		jsonResponse,_ := json.Marshal(err)
		w.WriteHeader(err.Code)
		w.Write(jsonResponse)
		return
	}
	jsonResponse,_ := json.Marshal(newUser)
	w.Write(jsonResponse)
}
