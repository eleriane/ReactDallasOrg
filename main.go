package main

import (
	"gopkg.in/kataras/iris.v6"
	//"gopkg.in/kataras/iris.v6/adaptors/view"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
)

func main() {
	app := iris.New()
	app.Adapt(httprouter.New())
	
	/*
	templates := view.HTML("./static", ".html")
    app.Adapt(templates)
	
	// Router
	app.Get("/", func(ctx *iris.Context) {
		ctx.Render("index.html", nil)
	})
	*/
	
	app.StaticWeb("/", "./static/")

	app.Listen(":80")
}