package utils

import (
	"encoding/json"
	"net/http"

	"github.com/mostafababaii/go-micro/services/users/config"
)

var (
	JsonResponse = &response{
		ContentType: config.JsonResponseContentType,
	}
)

type response struct {
	ContentType string
	StatusCode  int
}

func NewResponse(ct string, sc int) *response {
	return &response{
		ContentType: ct,
		StatusCode:  sc,
	}
}

func (r *response) Send(w http.ResponseWriter, content interface{}, params ...interface{}) {
	for _, param := range params {
		switch param.(type) {
		case int:
			r.StatusCode = param.(int)
		case string:
			r.ContentType = param.(string)
		}
	}

	w.Header().Set("Content-Type", r.ContentType)
	w.WriteHeader(r.StatusCode)

	switch content.(type) {
	case map[string]string, RestError:
		data, err := json.Marshal(content)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("error on marshaling response content"))
			return
		}

		w.Write(data)
		return
	case string:
		w.Write([]byte(content.(string)))
		return
	case []byte:
		w.Write(content.([]byte))
		return
	}

	w.Write([]byte("unable to send response"))
}
