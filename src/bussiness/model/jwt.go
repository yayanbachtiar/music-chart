package model

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	jwt.StandardClaims
	Phone string `json:"phone"`
	Name  string `json:"name"`
	Role  string `json:"role"`
}
