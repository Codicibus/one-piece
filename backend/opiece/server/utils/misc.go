package utils

import (
	"github.com/google/uuid"
	"opiece/server/middleware"
)

type Target string

func (t Target) In(x []string) bool {
	for _, val := range x {
		if t == Target(val) {
			return true
		}
	}
	return false
}

func GetBool(s string) bool {
	return s == "true" || s == "false"
}

func GetUUIDFromToken(token string) (uuid.UUID, error) {
	jwt := middleware.NewJWT()
	claims, err := jwt.ParseToken(token)
	if err != nil {
		return uuid.UUID{}, err
	}
	return claims.UUID, err
}
