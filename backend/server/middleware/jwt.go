package middleware

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"opiece/server/global"
	"opiece/server/model"
	"opiece/server/service/response"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("auth-token")
		if token == "" {
			response.FailWithDetailed(nil, "非法访问或未登录", c)
			c.Abort()
			return
		}
		// parse the JWT
		j := NewJWT()
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == TokenExpired {
				response.FailWithDetailed(nil, "token已过期", c)
				c.Abort()
				return
			}
			response.FailWithDetailed(nil, err.Error(), c)
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}

var (
	TokenExpired     = errors.New("token is expired")
	TokenNotValidYet = errors.New("token not active yet")
	TokenMalformed   = errors.New("that's not even a token")
	TokenInvalid     = errors.New("could not handle this token")
)

type JWT struct {
	SigningKey []byte
}

func NewJWT() *JWT {
	return &JWT{
		SigningKey: []byte(global.OConfig.JWT.SecretKey),
	}
}

func (j *JWT) CreateToken(claims model.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

func (j *JWT) ParseToken(tokenString string) (*model.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &model.CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*model.CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid
	}
	return nil, TokenInvalid
}
