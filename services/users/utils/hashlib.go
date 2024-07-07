package utils

import (
	"crypto/md5"
	"fmt"
)

var (
	Hashlib = &hashlib{}
)

type hashlib struct{}

func NewHashlib() *hashlib {
	return &hashlib{}
}

func (*hashlib) GetMD5(s string) string {
	h := md5.New()
	return fmt.Sprintf("%x", h.Sum([]byte(s)))
}
