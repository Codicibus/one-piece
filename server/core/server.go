package core

import (
	"opiece/server/global"
	"opiece/server/initialize"
	"opiece/server/middleware"
)

func RunServer() {
	global.OAuthJWT = middleware.NewJWTMiddleware()
	//db := initialize.NewGorm()
	router := initialize.Routers()
	//if db != nil {
		//initialize.MigrateTables(db)
		// TODO: close the db connection
	//}
	if err := router.Run(); err != nil {
		panic(err)
	}
}