package routers

import (
	"net/http"

	"github.com/mostafababaii/go-micro/services/users/app/handlers"
)

var (
	Default *http.ServeMux
)

func init() {
	sm := http.NewServeMux()
	sm.Handle("/users/", handlers.NewUserHandler())
	sm.Handle("/auth", handlers.NewAuthHandler())
	Default = sm
}
