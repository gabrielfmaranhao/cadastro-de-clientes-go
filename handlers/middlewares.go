package handlers

import (
	"cadastro_de_clientes/utils"
	"context"
	"encoding/json"
	"net/http"
	"strings"
)
func AuthenticateToken(next http.Handler) http.Handler  {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request)  {
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			err := &utils.HandlerError{
				Code: 400,
				Message: "Token is not exist or token invalid",
			}
			jsonResponse,_ := json.Marshal(err)
			w.WriteHeader(err.Code)
			w.Write(jsonResponse)
			return
		}
		token := strings.Split(authHeader, " ")[1]
		id,valid, err := VerifyJWTToken(token)
		if err != nil || !valid {
			err := &utils.HandlerError{
				Code: 400,
				Message: "Token is not exist or token invalid",
			}
			jsonResponse,_ := json.Marshal(err)
			w.WriteHeader(err.Code)
			w.Write(jsonResponse)
			return
		}
		ctx := context.WithValue(r.Context(), "id", id)
		next.ServeHTTP(w,r.WithContext(ctx))
	})
}
