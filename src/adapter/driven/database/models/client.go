package database_model

import "gorm.io/gorm"

type ClientModel struct {
	gorm.Model
	Id, Name, Cpf string
}
