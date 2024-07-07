package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	handler *Handler
)

type Config struct {
	User string
	Pass string
	Host string
	Port string
	Name string
}

type Handler struct {
	DB *gorm.DB
}

func NewHandler(config *Config) (*Handler, error) {
	if handler != nil {
		return handler, nil
	}

	dns := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.User,
		config.Pass,
		config.Host,
		config.Port,
		config.Name,
	)

	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Handler{DB: db}, nil
}
