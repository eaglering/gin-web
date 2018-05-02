package api

import (
	"github.com/gin-gonic/gin"
	"gin-web/app/modules/api/controllers"
	"gin-web/app/modules/api/middlewares"
)

func Routers(app *gin.Engine) {
	apiRouter := app.Group("/api")
	apiRouter.Use(middlewares.Env{}.New())

	c := controllers.Hello{}
	hello := apiRouter.Group("/hello")
	hello.GET("", c.Index)

}