package app

import (
	. "gin-web/app/config"
	"github.com/gin-gonic/gin"
	"gin-web/app/modules/api"
	"gin-web/app/common/pools/databases"
	"github.com/go-xorm/xorm"
	"gin-web/app/common/libraries/cache/persistence"
	"gin-web/app/common/pools/caches"
	"gin-web/app/common/pools/redis"
	redis2 "github.com/garyburd/redigo/redis"
	"log"
	"gin-web/app/modules/ws"
)

func setupRouter(app *gin.Engine) {
	api.Routers(app)
	ws.Routers(app)
}

func setupDatabase() *xorm.Engine {
	return databases.Instance()
}

func setupCache() persistence.CacheStore {
	return caches.Instance()
}

func setupRedis() *redis2.Pool {
	return redis.Instance()
}

func Bootstrap() {
	// 配置初始化
	if err := InitConfig(); err != nil {
		panic(err)
	}

	if Config.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
	app := gin.Default()

	setupRouter(app)

	db := setupDatabase()
	defer db.Close()
	log.Println(db)

	cache := setupCache()
	defer cache.Close()

	rs := setupRedis()
	defer rs.Close()

	//if Config.Server.Graceful {
	//	endless.ListenAndServe(Config.Server.Address, app)
	//} else {
		app.Run(Config.Server.Address)
	//}

}
