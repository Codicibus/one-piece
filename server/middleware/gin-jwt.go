package middleware

import (
	"opiece/server/model"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func payloadFunc(data interface{}) jwt.MapClaims {
	if v, ok := data.(*model.User); ok {
		return jwt.MapClaims{
			"id": v.Username,
		}
	}
	return jwt.MapClaims{}
}
func identityHandler(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	return &model.User{
		Username: claims["id"].(string),
	}
}
func authenticator(c *gin.Context) (interface{}, error) {
	var loginVales model.User
	if err := c.ShouldBind(&loginVales); err != nil {
		return "", jwt.ErrMissingLoginValues
	}
	userID := loginVales.Username
	password := loginVales.Password

	if (userID == "admin" && password == "admin") || (userID == "test" && password == "test") {
		return &model.User{
			Username: userID,
		}, nil
	}

	return nil, jwt.ErrFailedAuthentication
}
func authorizator(data interface{}, c *gin.Context) bool {
	if v, ok := data.(*model.User); ok && v.Username == "admin" {
		return true
	}

	return false
}
func unauthorized(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
}
func NewJWTMiddleware() *jwt.GinJWTMiddleware {
	config := &jwt.GinJWTMiddleware{
		Realm:           "test zone",
		Key:             []byte("secret key"),
		Timeout:         time.Hour,
		MaxRefresh:      time.Hour,
		IdentityKey:     "id",
		PayloadFunc:     payloadFunc,
		IdentityHandler: identityHandler,
		Authenticator:   authenticator,
		Authorizator:    authorizator,
		Unauthorized:    unauthorized,
		TokenLookup:     "header: Authorization, query: token, cookie: jwt",
		TokenHeadName:   "Bearer",
		TimeFunc:        time.Now,
	}
	j, err := jwt.New(config)
	if err != nil {
		panic(err)
	}
	err = j.MiddlewareInit()
	if err != nil {
		panic(err)
	}
	return j
}
