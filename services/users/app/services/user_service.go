package services

import (
	"net/http"
	"strconv"

	"github.com/mostafababaii/go-micro/services/users/app/models"
	"github.com/mostafababaii/go-micro/services/users/utils"
)

var (
	UserService = NewUserService()
)

type userService struct{}

func NewUserService() *userService {
	return &userService{}
}

func (s *userService) GetUserByID(id int) (*models.User, utils.RestError) {
	strID := strconv.Itoa(id)

	user := models.User{}
	user.FirstWhere("id", strID)

	err := models.Validate.Struct(user)
	if err != nil {
		return nil, utils.NewRestError(http.StatusNotFound, "user not found", "")
	}

	return &user, nil
}

func (s *userService) GetUserByCredentials(username, password string) (*models.User, utils.RestError) {
	user := models.User{}
	user.GetByUsernameAndPassword(username, password)

	err := models.Validate.Struct(user)
	if err != nil {
		return nil, utils.NewRestError(http.StatusNotFound, "user not found", "")
	}

	return &user, nil
}

func (s *userService) Create(user *models.User) (*models.User, utils.RestError) {
	err := models.Validate.Struct(user)
	if err != nil {
		return nil, utils.NewRestError(
			http.StatusBadRequest,
			"Error on validating",
			err.Error(),
		)
	}

	user.Password = utils.Hashlib.GetMD5(user.Password)

	affectedRows, err := user.Create()
	if err != nil || affectedRows == 0 {
		return nil, utils.NewRestError(
			http.StatusBadRequest,
			"registration process has not been successful",
			err.Error(),
		)
	}

	return user, nil
}
