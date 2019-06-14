package web

import "github.com/kataras/iris"

func apiAuthMiddleware(ctx iris.Context) {
	ctx.Text("api middleware")
	ctx.Next()
}
