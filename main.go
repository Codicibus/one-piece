package main

import (
	"opiece/server/core"
)

// @title OPiece
// @version 0.0.1
// @description 博客后台程序REST API接口
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /v1
func main() {
	core.RunServer()
}
