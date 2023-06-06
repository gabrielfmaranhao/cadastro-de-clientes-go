package handlers

import (
	"cadastro_de_clientes/models"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)
type createCellphone struct {
	Number string
}
func PostCellphone(w http.ResponseWriter, r *http.Request) {
	var cellphone createCellphone
	w.Header().Set("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&cellphone)
	idClient := chi.URLParam(r, "id")
	NewCellphone,err := models.NewCellphone(cellphone.Number,idClient)
	if err != nil {
		jsonResponse,_ := json.Marshal(err)
		w.WriteHeader(err.Code)
		w.Write(jsonResponse)
		return
	}
	jsonResponse,_ := json.Marshal(NewCellphone)
	w.WriteHeader(201)
	w.Write(jsonResponse)
}
func GetCellphones(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	cellphones,err := models.GetCellphones()
	if err != nil {
		jsonResponse,_ := json.Marshal(err)
		w.WriteHeader(err.Code)
		w.Write(jsonResponse)
		return
	}
	jsonResponse,_ := json.Marshal(cellphones)
	w.WriteHeader(200)
	w.Write(jsonResponse)
}
func DeleteCellphone(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	idCellphone := chi.URLParam(r, "id")
	err := models.DeleteCellPhone(idCellphone)
	if err != nil {
		jsonResponse,_ := json.Marshal(err)
		w.WriteHeader(err.Code)
		w.Write(jsonResponse)
		return
	}
	w.WriteHeader(204)
}
