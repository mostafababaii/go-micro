package utils

import (
	"net/url"
	"path"
)

var (
	Urllib = &urllib{}
)

type urllib struct{}

func NewUrllib() *urllib {
	return &urllib{}
}

func (*urllib) LastSegment(u string) (string, error) {
	url, err := url.Parse(u)
	if err != nil {
		return "", err
	}

	return path.Base(url.Path), nil
}
