package models

import (
	db "cadastro_de_clientes/config"
	handlerError "cadastro_de_clientes/utils"
	"fmt"
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)
type Cellphone struct{
	Id string `json:"id" gorm:"type:uuid;primary_key" valid:"uuid"`
	Number string `json:"number" gorm:"type:varchar(11);not null;unique" valid:"required"`
	CreatedAt time.Time `json:"created_at" valid:"-"`
	UpdatedAt time.Time `json:"updated_at" valid:"-"`
	ClientID string `json:"-" valid:"-"`
}
func NewCellphone(number, id string)(*Cellphone, *handlerError.HandlerError){
	cellphone := &Cellphone{
		Number: number,
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
	cellphone.Id = uuid.NewV4().String()
	cellphone.CreatedAt = time.Now()
	cellphone.ClientID = id
	err = cellphone.validate()
	if err != nil {
		return nil, &handlerError.HandlerError{
			Code: 400,
			Message: err.Error(),
		}
	}
	erro := conn.Create(cellphone).Scan(&cellphone).Error
	if erro != nil {
		return nil, &handlerError.HandlerError{
			Code: 400,
			Message: fmt.Sprintf("%v", erro),
		}
	}
	return cellphone, nil
}
func (cellphone *Cellphone)  validate() error {
	_, err := govalidator.ValidateStruct(cellphone)
	if err != nil {
		return err
	}
	return nil
}

func GetCellphones()([]Cellphone, *handlerError.HandlerError){
	var cellphones []Cellphone
	conn, err := db.OpenConnection()
	if err != nil {
		return nil, &handlerError.HandlerError{
			Code: 500,
			Message: fmt.Sprintf("Error prepare: %v", err),
		}
	}
	conn.Model(&cellphones).Preload("ClientID").Find(&cellphones)
	return cellphones,nil
}

func DeleteCellPhone(id string) *handlerError.HandlerError{
	var cellphone = Cellphone{
		Id: id,
	}
	conn, err := db.OpenConnection()
	if err != nil {
		return  &handlerError.HandlerError{
			Code: 500,
			Message: fmt.Sprintf("Error prepare: %v", err),
		}
	}
	err = conn.Where("id = ? ", id).First(&cellphone).Delete(&cellphone).Error
	if err != nil {
		return &handlerError.HandlerError{
			Code: 400,
			Message: err.Error(),
		}
	}
	return nil
}

func UpdateCelphone(number, id string) (*Cellphone, *handlerError.HandlerError){
	var cellphone Cellphone
	conn, err := db.OpenConnection()
	if err != nil {
		return nil, &handlerError.HandlerError{
			Code: 500,
			Message: fmt.Sprintf("Error prepare: %v", err),
		}
	}
	_ = conn.Where("number = ?", number).First(&cellphone).Error
	if  cellphone.Id != "" {
		return nil, &handlerError.HandlerError{
			Code: 400,
			Message: "Number is exists",
		}
	}
	err = conn.Where("id = ? ", id).First(&cellphone).Error
	if err != nil {
		return nil, &handlerError.HandlerError{
			Code: 400,
			Message: "Not found",
		}
	}
	cellphone.Number = number
	err = conn.Model(&cellphone).Updates(cellphone).Error
	if err != nil {
		return nil,&handlerError.HandlerError{
			Code: 400,
			Message: err.Error(),
		}
	}
	return &cellphone,nil

}
