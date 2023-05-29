package models

import (
	db "cadastro_de_clientes/config"
	handlerError "cadastro_de_clientes/utils"
	"fmt"
	"time"
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	jwt "github.com/dgrijalva/jwt-go"
	"os"
	"github.com/joho/godotenv"
)

type User struct {
	Id string `json:"id" gorm:"type:uuid;primary_key" valid:"uuid"`
	Username string `json:"username" gorm:"type:varchar(255);unique" valid:"required"`
	Cpf string `json:"cpf" gorm:"type:varchar(11);unique" valid:"required"`
	Password string `json:"-" gorm:"type:varchar(255)" valid:"notnull"`
	CreatedAt time.Time `json:"created_at" valid:"-"`
	UpdatedAt time.Time `json:"updated_at" valid:"-"`
}
func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}
func NewUser(username, cpf, password  string) (*User, *handlerError.HandlerError) {
	user := &User{
		Username:     username,
		Cpf:    cpf,
		Password: password,
	}
	errou := handlerError.ValidateCpf(user.Cpf)
	if errou != nil {
		return nil,errou
	}
	err := user.Prepare()
	if err != nil {
		return nil, &handlerError.HandlerError{
			Code: 400,
			Message: fmt.Sprintf("Error prepare: %v", err),
		}
	}
	conn, err := db.OpenConnection()
	if err != nil {
		return nil, &handlerError.HandlerError{
			Code: 500,
			Message: fmt.Sprintf("Error conect bank: %v", err),
		}
	}
	erro := conn.Create(user).Scan(&user)
	if erro.Error != nil {
		return nil, &handlerError.HandlerError{
			Message: fmt.Sprintf("%v", erro.Error),
			Code: 400,
		}
	}
	return user, nil
}

func (user *User) Prepare() error {

	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}
	user.Id = uuid.NewV4().String()
	user.CreatedAt = time.Now()
	user.Password = string(password)
	err = user.validate()
	if err != nil {
		return err
	}
	return nil
}

func (user *User) IsCorrectPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}

func (user *User) validate() error {

	_, err := govalidator.ValidateStruct(user)

	if err != nil {
		return err
	}
	return nil
}

func LoginUser(username, password string)(string, *handlerError.HandlerError) {
	err := godotenv.Load()
	if err != nil {
		return "",&handlerError.HandlerError{
			Code: 500,
			Message: err.Error(),
		}
	}
	var user User
	conn, err := db.OpenConnection()
	if err != nil {
		return "", &handlerError.HandlerError{
			Code: 500,
			Message: err.Error(),
		}
	}
	erro := conn.Where("username = ?", username).First(&user).Scan(&user)
	if erro.Error != nil {
		return "", &handlerError.HandlerError{
			Code: 400,
			Message: erro.Error.Error(),
		}
	}
	if !user.IsCorrectPassword(password) {
		return "", &handlerError.HandlerError{
			Code: 400,
			Message: "Username or password incorrect",
		}
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = user.Id
	claims["user"] = user.Username
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", &handlerError.HandlerError{
			Code: 500,
			Message: err.Error(),
		}
	}
	return tokenString, nil
}

