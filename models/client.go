package models

import (
	db "cadastro_de_clientes/config"
	handlerError "cadastro_de_clientes/utils"
	"time"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"github.com/asaskevich/govalidator"
)
type Client struct{
	Id string `json:"id" gorm:"type:uuid;primary_key" valid:"uuid"`
	Name string `json:"name" gorm:"type:varchar(255);not null" valid:"required"`
	Cpf string `json:"cpf" gorm:"type:varchar(11);unique" valid:"required"`
	User_at string `json:"user_at" gorm:"type:varchar(255)" valid:"-"`
	CreatedAt time.Time `json:"created_at" valid:"-"`
	UpdatedAt time.Time `json:"updated_at" valid:"-"`
	UserID string `json:"-" valid:"-"`
	Cellphones []Cellphone `json:"cellphones" valid:"-"`
	Emails []Email `json:"emails" valid:"-"`
}

func GetClients() ([]Client, *handlerError.HandlerError) {
	var clients []Client
	conn, err := db.OpenConnection()
	if err != nil {
		return clients, &handlerError.HandlerError{
			Code: 500,
			Message: fmt.Sprintf("Error prepare: %v", err),
		}
	}
	conn.Model(&clients).Preload("Emails").Preload("Cellphones").Find(&clients)
	return clients,nil
}
func NewCLient(name, cpf, id string) (*Client, *handlerError.HandlerError){
	 client := &Client{
		Name: name,
	 }
	var user User
	conn, err := db.OpenConnection()
	if err != nil {
		return nil, &handlerError.HandlerError{
			Code: 500,
			Message: fmt.Sprintf("Error prepare: %v", err),
		}
	}
	conn.Where("cpf = ?", cpf).First(&user).Scan(&user)
	if user.Id != ""{
		return nil, &handlerError.HandlerError{
			Code: 400,
			Message: "Cpf is exist",
		}
	}
	user, errorPerson := Profile(id)
	if errorPerson != nil {
		return nil, errorPerson
	}
	errorPerson = handlerError.ValidateCpf(cpf)
	if errorPerson != nil {
		return nil, errorPerson
	}
	client.Id = uuid.NewV4().String()
	client.Cpf = cpf
	client.User_at = user.Username
	client.CreatedAt = time.Now()
	client.UserID = user.Id
	err = client.validate()
	if err != nil {
		return nil, &handlerError.HandlerError{
			Code: 400,
			Message: err.Error(),
		}
	}
	erro := conn.Create(client).Scan(&client)
	if erro.Error != nil {
		return nil, &handlerError.HandlerError{
			Code: 400,
			Message: fmt.Sprintf("%v", erro.Error),
		}
	}
	fmt.Println(client)
	return client, nil
}
func (client *Client) validate() error{
	_, err := govalidator.ValidateStruct(client)
	if err != nil {
		return err
	}
	return nil
}
func ClientProfile(id string) (*Client, *handlerError.HandlerError){
	var client Client
	conn, err := db.OpenConnection()
	if err != nil {
		return nil, &handlerError.HandlerError{
			Code: 500,
			Message: fmt.Sprintf("Error prepare: %v", err),
		}
	}
	// erro := conn.Where("id = ?", id).Preload("Clients.Emails").Preload("Clients.Cellphones").First(&user)
	erro := conn.Model(client).Preload("Emails").Preload("Cellphones").Where("id = ?", id).First(&client)
	if erro.Error != nil {
		return nil , &handlerError.HandlerError{
			Code: 400,
			Message: erro.Error.Error(),
		}
	}
	return &client, nil
}
