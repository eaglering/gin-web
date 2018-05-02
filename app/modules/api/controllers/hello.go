package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Hello struct{}

func (c *Hello) Index(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"msg": "hello",
	})
}