package handlers

import (
	"os"
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)
func VerifyJWTToken(tokenString string) (interface{} ,bool, error) {
	err := godotenv.Load()
		if err != nil  {
			return "",false, err
		}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return []byte(os.Getenv("SECRET_KEY")), nil
    })
	if err != nil {
        return "", false, err
    }
	if !token.Valid {
        return "", false, nil
    }
	claims := token.Claims.(jwt.MapClaims)
	id := claims["sub"]
	return id, true, nil
}
