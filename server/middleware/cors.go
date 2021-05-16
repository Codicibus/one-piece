package middleware

import "github.com/gin-contrib/cors"

func NewCorsConfig() cors.Config {
	return cors.DefaultConfig()
}
