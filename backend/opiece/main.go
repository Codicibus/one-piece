package main

import (
	"opiece/server/core"
)

// @title OPiece
// @version 0.0.1
// @description 博客后台程序REST API接口
// @termsOfService http://swagger.io/terms/

// @contact.name aumujun@gmail.com
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host api.amujun.com
// @BasePath /v1
func main() {
	core.RunServer()
}
