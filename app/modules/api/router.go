package api

import (
	"github.com/gin-gonic/gin"
	"gin-web/app/modules/api/controllers"
	"gin-web/app/modules/api/middlewares"
)

func Routers(router *gin.Engine) {
	api := router.Group("/api")
	{
		env := &middlewares.Env{}
		any := api.Group("/any", env.New())
		{
			c := &controllers.Hello{}
			any.GET("/hello", c.Index)
			any.GET("/database", c.TestDb)
			any.GET("/cache", c.TestCache)
			any.GET("/redis", c.TestRedis)
		}

		safe := api.Group("/safe", env.New(), env.Authorize())
		{
			c := &controllers.Hello{}
			safe.GET("/hello", c.Index)
		}
	}

}