package models

import (
	"github.com/go-playground/validator/v10"
)

var (
	Validate *validator.Validate
)

func init() {
	Validate = validator.New()
}
