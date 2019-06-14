package web

import (
	"fmt"
	"github.com/indes/flowerss-bot/config"
	"github.com/kataras/iris"
)

var (
	port     = config.WebPort
	botToken = config.BotToken
)

func init() {

}

func Run() {
	app := iris.Default()
	routeInit(app)
	app.RegisterView(iris.HTML("./web/views", ".html"))
	// listen and serve on http://0.0.0.0:8080.

	app.Run(iris.Addr(fmt.Sprintf(":%d", port)))
}
