package handlers

import (
	"cadastro_de_clientes/models"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)
type createEmail struct {
	Email string
}
func PostEmail(w http.ResponseWriter, r *http.Request) {
	var email createEmail
	w.Header().Set("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&email)
	idClient := chi.URLParam(r, "id")
	newEmail,err := models.NewEmail(email.Email,idClient)
	if err != nil {
		jsonResponse,_ := json.Marshal(err)
		w.WriteHeader(err.Code)
		w.Write(jsonResponse)
		return
	}
	jsonResponse,_ := json.Marshal(newEmail)
	w.WriteHeader(201)
	w.Write(jsonResponse)
}
func GetEmails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	emails,err := models.Emails()
	if err != nil {
		jsonResponse,_ := json.Marshal(err)
		w.WriteHeader(err.Code)
		w.Write(jsonResponse)
		return
	}
	jsonResponse,_ := json.Marshal(emails)
	w.WriteHeader(200)
	w.Write(jsonResponse)
}
func DeleteEmail(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	idEmail := chi.URLParam(r, "id")
	err := models.DeleteEmail(idEmail)
	if err != nil {
		jsonResponse,_ := json.Marshal(err)
		w.WriteHeader(err.Code)
		w.Write(jsonResponse)
		return
	}
	w.WriteHeader(204)
}
func UpdateEmail(w http.ResponseWriter, r *http.Request){
	var email createEmail
	w.Header().Set("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&email)
	idEmail := chi.URLParam(r, "id")
	response,err := models.UpdateEmail(email.Email, idEmail)
	if err != nil {
		jsonResponse,_ := json.Marshal(err)
		w.WriteHeader(err.Code)
		w.Write(jsonResponse)
		return
	}
	jsonResponse,_ := json.Marshal(response)
	w.WriteHeader(201)
	w.Write(jsonResponse)
}
