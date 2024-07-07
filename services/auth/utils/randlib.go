package utils

import (
	"math/rand"
	"time"
)

var (
	Randlib     = &randlib{}
	letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

type randlib struct{}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func NewRandlib() *randlib {
	return &randlib{}
}

func (*randlib) RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
