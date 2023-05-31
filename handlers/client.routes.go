package handlers

import (
	"cadastro_de_clientes/models"
	"encoding/json"
	"net/http"
)
type create struct {
	Name string
	Cpf string
}
func Get(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	clients,err := models.GetClients()
	if err != nil {
		jsonResponse,_ := json.Marshal(err)
		w.WriteHeader(err.Code)
		w.Write(jsonResponse)
		return
	}
	jsonResponse,_ := json.Marshal(clients)
	w.Write(jsonResponse)
}
func Post(w http.ResponseWriter, r *http.Request) {
	var client create
	id := r.Context().Value("id").(string)
	w.Header().Set("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&client)
	newClient, err := models.NewCLient(client.Name, client.Cpf, id)
	if err != nil {
		jsonResponse,_ := json.Marshal(err)
		w.WriteHeader(err.Code)
		w.Write(jsonResponse)
		return
	}
	jsonResponse,_ := json.Marshal(newClient)
	w.Write(jsonResponse)
}
