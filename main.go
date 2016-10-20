package main

import (
	//"github.com/iris-contrib/middleware/logger"
	"github.com/iris-contrib/middleware/recovery"
	"github.com/kataras/iris"
)

func main() {

	// Middleware
	iris.Use(recovery.New())
	//iris.Use(logger.New())
	
	iris.StaticFS("/s", "./s/", 1)
	//iris.StaticWeb("/","./static", 1)

	// Router
	iris.Get("/", func(ctx *iris.Context) {
		ctx.ServeFile("./s/index.html", true)
	})

	iris.Listen(":80")
}