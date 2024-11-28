package models

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"Email" gorm:"unique" validate:"required,email"`
	Password  string `json:"Password" validate:"required,min=8"`
}

var Validate *validator.Validate

func init() {
	Validate = validator.New()
}
