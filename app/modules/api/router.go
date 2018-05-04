package api

import (
	"github.com/gin-gonic/gin"
	"gin-web/app/modules/api/controllers"
)

func Routers(router *gin.Engine) {
	api := router.Group("/api")
	{
		any := api.Group("/any")
		{
			c := &controllers.Hello{}
			any.GET("/hello", c.Index)
			any.GET("/database", c.TestDb)
		}
		//env := &middlewares.Env{}
		//safe := api.Group("/safe", env.New())
		//{
		//	c := &controllers.Hello{}
		//	safe.GET("/hello", c.Index)
		//}
	}

}