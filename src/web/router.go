package web

import (
	"github.com/iissy/goweb/src/controller"
	"github.com/iissy/goweb/src/middleware"
	"github.com/kataras/iris/v12"
)

func Start() {
	app := iris.New()
	app.Use(iris.LimitRequestBodySize(1024 * 1024))
	app.HandleDir("/", "./public")
	tmpl := iris.HTML("./views", ".html")
	tmpl.Layout("shared/layout.html")
	tmpl.Reload(true)
	app.RegisterView(tmpl)
	app.Use(middleware.TraceWeb)

	app.Get("/", controller.Index)
	app.Get("/links/{id}", controller.ListLinks)
	app.Get("/cuslinks", controller.ListCusLinks)
	app.Get("/articles", controller.ListArticles)
	app.Get("/article/{id}", controller.Detail)
	app.Get("/payme", controller.Payme)
	app.Get("/link/{id}", controller.GetLinkUrl)
	app.Get("/cuslink/{id}", controller.GetCusLinkUrl)
	app.Get("/choose", controller.Choose)
	//app.Post("/mosaic", controller.Mosaic)

	app.Run(
		iris.Addr(":80"),
		iris.WithoutBanner,
		iris.WithoutServerError(iris.ErrServerClosed),
	)
}
