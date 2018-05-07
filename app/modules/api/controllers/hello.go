package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"gin-web/app/common/pools/databases"
	"log"
	"gin-web/app/common/models"
	"gin-web/app/common/helpers"
	"gin-web/app/common/pools/caches"
	"gin-web/app/common/pools/redis"
	redis2 "github.com/garyburd/redigo/redis"
)

type Hello struct{}

func (c *Hello) Index(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, helpers.JSONFormat("hello", gin.H{
		"name": "eaglering",
	}))
}

/**
 * Firstly, install xorm command line:
 * 	  go get github.com/go-xorm/cmd
 * Then, generate codes for Go:
 *    xorm reverse -s mysql root:password@tcp(host:port)/database?charset=utf8 \
 *        src/github.com/go-xorm/cmd/xorm/templates/goxorm \
 *        src/gin-web/app/common/models
 */
func (c *Hello) TestDb(ctx *gin.Context) {
	db := databases.Instance()
	log.Println(db)
	user := &models.Users{
		Id: 1,
		Name: "eaglering",
	}
	_, err := db.InsertOne(user)
	if err != nil {
		log.Println(err)
	}
	u := &models.Users{
		Id: 1,
	}
	has, err := db.Get(u)
	if err != nil {
		log.Println(err)
	}
	log.Println(has)
	ctx.JSON(http.StatusOK, helpers.JSONFormat("", u))
}

func (c *Hello) TestCache(ctx *gin.Context) {
	cache := caches.Instance()
	err := cache.Set("test", "" , caches.Forever)
	if err != nil {
		log.Println(err)
	}
	var result interface{}
	err = cache.Get("test", &result)
	if err != nil {
		log.Println(err)
	}
	ctx.JSON(http.StatusOK, helpers.JSONFormat("", result))
}

func (c *Hello) TestRedis(ctx *gin.Context) {
	r := redis.Instance()
	conn := r.Get()
	defer conn.Close()
	result, err := redis2.String(conn.Do("GET", "test"))
	if err != nil {
		log.Println(err)
	}
	ctx.JSON(http.StatusOK, helpers.JSONFormat("", result))
}