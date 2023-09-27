package common

import "github.com/dgrijalva/jwt-go"

type UserJwt struct {
	Username string
	UserId   uint
	UserType uint
	jwt.StandardClaims
}
