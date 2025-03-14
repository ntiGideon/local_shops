package model

import "github.com/dgrijalva/jwt-go"

type JWTPayload struct {
	Email string `json:"email"`
	Id    int    `json:"id"`
	Role  string `json:"role"`
}

type JWTClaim struct {
	Id         int         `json:"id"`
	Role       string      `json:"role"`
	Email      string      `json:"email"`
	JwtId      string      `json:"jwtId"`
	CustomData interface{} `json:"customData"`
	jwt.StandardClaims
}
