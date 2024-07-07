package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/mostafababaii/go-micro/services/users/app/services"
	"github.com/mostafababaii/go-micro/services/users/utils"
)

type authHandler struct{}

func NewAuthHandler() *authHandler {
	return &authHandler{}
}

func (l *authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		l.auth(w, r)
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (*authHandler) auth(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.JsonResponse.Send(w, err.Error(), http.StatusBadRequest)
	}

	auth := services.NewAuthRequest()
	if err = json.Unmarshal(data, auth); err != nil {
		utils.JsonResponse.Send(w, err.Error(), http.StatusBadRequest)
	}

	user, re := services.UserService.GetUserByCredentials(
		auth.Username,
		auth.Password,
	)

	if re != nil {
		utils.JsonResponse.Send(w, re, re.Status())
		return
	}

	token, err := services.AuthService.GetToken(user.ID)
	if err != nil {
		panic(err)
	}

	utils.JsonResponse.Send(w, map[string]string{"token": token}, http.StatusOK)
}
