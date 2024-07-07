package domain

import (
	"errors"
	"os/exec"
	"strings"
	"time"

	"github.com/mostafababaii/go-micro/services/auth/utils"
)

type Token struct {
	ID          string
	AccessToken string
	UserId      uint
	Expires     time.Time
}

func NewToken(userID uint) (*Token, error) {
	uuid, err := exec.Command("uuidgen").Output()
	if err != nil {
		return nil, errors.New("uuid generation has been failed")
	}
	oneMonth := time.Hour * time.Duration(720)
	token := &Token{
		ID:          strings.Trim(string(uuid), "\n"),
		AccessToken: utils.Randlib.RandStringRunes(64),
		UserId:      userID,
		Expires:     time.Now().Local().Add(oneMonth),
	}
	return token, nil
}
