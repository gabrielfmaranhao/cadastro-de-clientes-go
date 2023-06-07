package models
import (
	db "cadastro_de_clientes/config"
	handlerError "cadastro_de_clientes/utils"
	"time"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"github.com/asaskevich/govalidator"
)
type Email struct{
	Id string `json:"id" gorm:"type:uuid;primary_key" valid:"uuid"`
	Email string `json:"email" gorm:"type:varchar(255);not null;unique" valid:"required"`
	CreatedAt time.Time `json:"created_at" valid:"-"`
	UpdatedAt time.Time `json:"updated_at" valid:"-"`
	ClientID string `json:"client" valid:"-"`
}
func NewEmail(email, id string)(*Email, *handlerError.HandlerError){
	new := &Email{
		Email: email,
	}
	var client Client
	conn, err := db.OpenConnection()
	if err != nil {
		return nil, &handlerError.HandlerError{
			Code: 500,
			Message: fmt.Sprintf("Error prepare: %v", err),
		}
	}
	conn.Where("id = ?", id).First(&client)
	new.Id = uuid.NewV4().String()
	new.CreatedAt = time.Now()
	new.ClientID = id
	err = new.validate()
	if err != nil {
		return nil, &handlerError.HandlerError{
			Code: 400,
			Message: err.Error(),
		}
	}
	erro := conn.Create(new).Scan(&new)
	if erro.Error != nil {
		return nil, &handlerError.HandlerError{
			Code: 400,
			Message: fmt.Sprintf("%v", erro.Error),
		}
	}
	return new, nil
}
func (email *Email)  validate() error {
	_, err := govalidator.ValidateStruct(email)
	if err != nil {
		return err
	}
	return nil
}
func Emails()([]Email, *handlerError.HandlerError){
	var emails []Email
	conn, err := db.OpenConnection()
	if err != nil {
		return nil, &handlerError.HandlerError{
			Code: 500,
			Message: fmt.Sprintf("Error prepare: %v", err),
		}
	}
	conn.Model(&emails).Preload("ClientID").Find(&emails)
	return emails,nil
}
func DeleteEmail(id string) *handlerError.HandlerError{
	var email = Email{
		Id: id,
	}
	conn, err := db.OpenConnection()
	if err != nil {
		return  &handlerError.HandlerError{
			Code: 500,
			Message: fmt.Sprintf("Error prepare: %v", err),
		}
	}
	err = conn.Where("id = ? ", id).First(&email).Delete(&email).Error
	if err != nil {
		return &handlerError.HandlerError{
			Code: 400,
			Message: err.Error(),
		}
	}
	return nil
}
func UpdateEmail(emailString, id string) (*Email, *handlerError.HandlerError){
	var email Email
	conn, err := db.OpenConnection()
	if err != nil {
		return nil, &handlerError.HandlerError{
			Code: 500,
			Message: fmt.Sprintf("Error prepare: %v", err),
		}
	}
	_ = conn.Where("email = ?", emailString).First(&email).Error
	if  email.Id != "" {
		return nil, &handlerError.HandlerError{
			Code: 400,
			Message: "Email is exists",
		}
	}
	err = conn.Where("id = ? ", id).First(&email).Error
	if err != nil {
		return nil, &handlerError.HandlerError{
			Code: 400,
			Message: "Not found",
		}
	}
	email.Email = emailString
	err = conn.Model(&email).Updates(email).Error
	if err != nil {
		return nil,&handlerError.HandlerError{
			Code: 400,
			Message: err.Error(),
		}
	}
	return &email,nil

}
