package handlers

import (
	"cadastro_de_clientes/models"
	"encoding/json"
	"net/http"
)
type register struct {
	Username string
	Cpf string
	Password string
}
type login struct {
	Username string
	Password string
}
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user register
	w.Header().Set("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&user)
	newUser,err := models.NewUser(user.Username, user.Cpf, user.Password)
	if err != nil {
		jsonResponse,_ := json.Marshal(err)
		w.WriteHeader(err.Code)
		w.Write(jsonResponse)
		return
	}
	jsonResponse,_ := json.Marshal(newUser)
	w.Write(jsonResponse)
}
func LoginUser(w http.ResponseWriter, r *http.Request) {
	var user login
	w.Header().Set("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&user)
	token, err := models.LoginUser(user.Username, user.Password)
	if err != nil{
		jsonResponse,_ := json.Marshal(err)
		w.WriteHeader(err.Code)
		w.Write(jsonResponse)
		return
	}
	jsonResponse,_ := json.Marshal(token)
	w.Write(jsonResponse)
}
func ProfileUser(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("id").(string)
	var user register
	w.Header().Set("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&user)
	userUpdate,err := models.Profile(id)
	if err != nil {
		jsonResponse,_ := json.Marshal(err)
		w.WriteHeader(err.Code)
		w.Write(jsonResponse)
		return
	}
	jsonResponse,_ := json.Marshal(userUpdate)
	w.Write(jsonResponse)
}

