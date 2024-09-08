package models

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	Login string `json:"login"`
	jwt.RegisteredClaims
}

func (c *Claims) Valid() error {
	panic("unimplemented")
}
