package model

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type CustomClaims struct {
	UUID       uuid.UUID
	ID         uint
	Username   string
	BufferTime int64
	jwt.StandardClaims
}
