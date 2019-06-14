package web

import "github.com/kataras/iris"

func routeInit(app *iris.Application) {
	//_ = app.Get("/ping", func(ctx iris.Context) {
	//	_, _ = ctx.JSON(iris.Map{
	//		"message": "pong",
	//	})
	//})
	app.PartyFunc("/api", func(api iris.Party) {
		api.Use(apiAuthMiddleware)
		api.Get("/{id:int}/profile", func(ctx iris.Context) {
			ctx.Text("profile")
		})

		//api.Get("/inbox/{id:int}", userMessageHandler)
	})
	app.Get("/", HomeController)
	app.Get("/login", func(context iris.Context) {
		context.View("login.html")
	})
	app.Get("/auth", auth)
}
