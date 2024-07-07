package utils

import (
	"encoding/json"
	"net/http"
)

var (
	Jsonlib = &jsonlib{}
)

type jsonlib struct{}

func NewJsonlib() *jsonlib {
	return &jsonlib{}
}

func (*jsonlib) BindJson(data []byte, i interface{}) *restError {
	if err := json.Unmarshal(data, i); err != nil {
		return NewRestError(
			http.StatusBadRequest,
			"fail to unmarshal data",
			err.Error(),
		)
	}
	return nil
}
