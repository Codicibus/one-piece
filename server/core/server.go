package core

import (
	"opiece/server/config"
	"opiece/server/global"
	"opiece/server/initialize"
	"opiece/server/middleware"
)

func RunServer() {
	global.OConfig = config.ReadYAMLConfig()
	global.OAuthJWT = middleware.NewJWTMiddleware()
	global.ODB = initialize.NewGorm()
	router := initialize.Routers()
	if global.ODB != nil {
		initialize.MigrateTables(global.ODB)
		// TODO: close the db connection
		db,_ := global.ODB.DB()
		defer db.Close()
	}
	if err := router.Run(); err != nil {
		panic(err)
	}
}