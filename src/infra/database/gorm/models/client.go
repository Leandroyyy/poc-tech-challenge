package database_model

import "gorm.io/gorm"

type ClientModel struct {
	gorm.Model
	Name, Cpf string
}
