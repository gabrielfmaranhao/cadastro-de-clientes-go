package models

import (
	"cadastro_de_clientes/config"
)

func Migrate() error {
	conn, err := config.OpenConnection()
	if err != nil {
		return err
	}
	conn.AutoMigrate(User{}, Client{})
	return nil
}
