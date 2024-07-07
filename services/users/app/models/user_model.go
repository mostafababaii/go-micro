package models

import (
	"fmt"
	"log"

	"github.com/mostafababaii/go-micro/services/users/config"
	"github.com/mostafababaii/go-micro/services/users/db"
	"github.com/mostafababaii/go-micro/services/users/utils"
	"gorm.io/gorm"
)

var (
	handler *db.Handler
)

func init() {
	var err error
	handler, err = db.NewHandler(config.DatabaseConfig)
	if err != nil {
		log.Fatal(err)
	}
	handler.DB.Migrator().AutoMigrate(&User{})
}

type User struct {
	gorm.Model `json:"-"`
	FirstName  string `gorm:"size:30" json:"first_name"`
	LastName   string `gorm:"size:30" json:"last_name"`
	Username   string `gorm:"size:30;index:idx_username,unique" json:"username" validate:"required"`
	Password   string `gorm:"size:50" json:"password" validate:"required"`
	Email      string `gorm:"size:100;index:idx_email,unique" json:"email" validate:"required"`
}

func (user *User) FirstWhere(key, value string) {
	handler.DB.First(user, fmt.Sprintf("`%s` = '%s'", key, value))
}

func (user *User) GetByUsernameAndPassword(username, password string) {
	handler.DB.Where(
		"`username` = ? AND `password` = ?",
		username,
		utils.Hashlib.GetMD5(password),
	).Find(user)
}

func (user *User) GetByUsername(username string) {
	user.FirstWhere("username", username)
}

func (user *User) GetByEmail(email string) {
	user.FirstWhere("email", email)
}

func (user *User) Create() (int64, error) {
	result := handler.DB.Create(user)
	return result.RowsAffected, result.Error
}
