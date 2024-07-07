package handlers

import (
	"net/http"
)

type ping struct{}

func NewPing() *ping {
	return &ping{}
}

func (*ping) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Pong"))
}
