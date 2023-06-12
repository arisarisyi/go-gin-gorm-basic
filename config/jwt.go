package config

import (
	"os"

	"github.com/golang-jwt/jwt/v4"
)

var JWT_SECRET_KEY = os.Getenv("JWT_SECRET_KEY")

var JWT_KEY = []byte(JWT_SECRET_KEY)

type JWTClaim struct {
	Email string
	Role string
	jwt.RegisteredClaims
}