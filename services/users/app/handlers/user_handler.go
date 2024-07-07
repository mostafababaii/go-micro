package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"

	"github.com/mostafababaii/go-micro/services/users/app/models"
	"github.com/mostafababaii/go-micro/services/users/app/services"
	"github.com/mostafababaii/go-micro/services/users/utils"
)

type userHandler struct{}

func NewUserHandler() *userHandler {
	return &userHandler{}
}

func (u *userHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		u.register(w, r)
		return
	}

	if r.Method == http.MethodGet {
		lastSegment, err := utils.Urllib.LastSegment(r.URL.RequestURI())
		if err != nil {
			utils.JsonResponse.Send(w, err.Error(), http.StatusInternalServerError)
			return
		}

		re := regexp.MustCompile("^[0-9]*$")
		matches := re.FindAllString(lastSegment, -1)

		if len(matches) > 0 {
			id, err := strconv.Atoi(matches[0])
			if err != nil {
				utils.JsonResponse.Send(w, err.Error(), http.StatusInternalServerError)
				return
			}
			u.info(w, id)
			return
		}

		utils.JsonResponse.Send(w, "url not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (*userHandler) register(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.JsonResponse.Send(w, err.Error(), http.StatusBadRequest)
	}

	user := &models.User{}

	if rerr := utils.Jsonlib.BindJson(data, &user); err != nil {
		utils.JsonResponse.Send(w, rerr, http.StatusBadRequest)
		return
	}

	u, re := services.UserService.Create(user)
	if re != nil {
		utils.JsonResponse.Send(w, re, re.Status())
		return
	}

	message := map[string]string{
		"message": fmt.Sprintf(
			"user has been successfully registered with email: %s",
			u.Email,
		),
	}

	utils.JsonResponse.Send(w, message, http.StatusCreated)
}

func (*userHandler) info(w http.ResponseWriter, id int) {
	u, re := services.UserService.GetUserByID(id)
	if re != nil {
		utils.JsonResponse.Send(w, re, re.Status())
		return
	}

	resp, err := json.Marshal(u)
	if err != nil {
		utils.JsonResponse.Send(w, utils.NewRestError(
			http.StatusInternalServerError,
			"error on marshaling content",
			err.Error(),
		))
		return
	}

	utils.JsonResponse.Send(w, resp, http.StatusOK)
}
