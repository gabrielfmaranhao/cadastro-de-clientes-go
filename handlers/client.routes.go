package handlers

import (
	"cadastro_de_clientes/models"
	"encoding/json"
	"net/http"
	"github.com/go-chi/chi"
)
// [] Update Client


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
	var client models.CreateClient
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
	models.NewEmail(client.Email, newClient.Id)
	models.NewCellphone(client.Number, newClient.Id)
	clientNew,_ := models.ClientProfile(newClient.Id)
	jsonResponse,_ := json.Marshal(clientNew)
	w.WriteHeader(201)
	w.Write(jsonResponse)
}
func Update(w http.ResponseWriter, r *http.Request) {
	var client models.CreateClient
	idClient := chi.URLParam(r, "id")
	w.Header().Set("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&client)
	update, err := models.UpdateClient(client, idClient)
	if err != nil  {
		jsonResponse,_ := json.Marshal(err)
		w.WriteHeader(err.Code)
		w.Write(jsonResponse)
		return
	}
	jsonResponse,_ := json.Marshal(update)
	w.WriteHeader(201)
	w.Write(jsonResponse)
}
func Delete(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	idClient := chi.URLParam(r, "id")
	err := models.DeleteClient(idClient)
	if err != nil {
		jsonResponse,_ := json.Marshal(err)
		w.WriteHeader(err.Code)
		w.Write(jsonResponse)
		return
	}
	w.WriteHeader(204)
}
