package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"gin-web/app/common/pools/databases"
	"log"
	"gin-web/app/common/models"
	"gin-web/app/common/helpers"
)

type Hello struct{}

func (c *Hello) Index(context *gin.Context) {
	context.JSON(http.StatusOK, helper.JSONFormat("hello", gin.H{
		"name": "eaglering",
	}))
}

/**
 * Firstly, install xorm command line:
 * 	  go get github.com/go-xorm/cmd
 * Then, generate codes for Go:
 *    xorm reverse -s mysql root:password@tcp(host:port)/database?charset=utf8 \
 *        src/github.com/go-xorm/cmd/xorm/templates/goxorm/struct.go.tpl \
 *        src/gin-web/app/common/models
 */
func (c *Hello) TestDb(context *gin.Context) {
	db := databases.Instance()
	log.Println(db)
	user := &models.DUsers{
		Id: 1,
		Mobile: "12345678901",
		Nickname: "eaglering",
		Sex: models.Mail,
		Country: "中国",
		Language: "zh-CN",
	}
	_, err := db.InsertOne(user)
	if err != nil {
		log.Println(err)
	}
	u := &models.DUsers{
		Id: 1,
	}
	has, err := db.Get(u)
	if err != nil {
		log.Println(err)
	}
	log.Println(has)
	context.JSON(http.StatusOK, helper.JSONFormat("", u))
}