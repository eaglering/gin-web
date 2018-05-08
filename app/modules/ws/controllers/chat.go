package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/lunny/log"
	"encoding/json"
	"gin-web/app/modules/ws/middlewares"
	"gin-web/app/common/helpers"
)

type Chat struct{}

func (c *Chat) Person(ctx *gin.Context) {
	s, _ := ctx.Get("sender")
	sender := s.(middlewares.Client)
	recipient, has := ctx.Params.Get("recipient")
	if !has || recipient == "" {
		sender.Send <- []byte(helpers.ERROR)
	}
	content, has := ctx.Params.Get("content")
	if !has || content == "" {
		sender.Send <- []byte(helpers.ERROR)
	}
	// todo
	m, _ := ctx.Get("manager")
	manager := m.(middlewares.Manager)
	for client, _ := range manager.Clients {
		// Online
		if client.User["uid"] == recipient {
			// todo
			log.Println(sender.User)
			log.Println(content)
			log.Println(client)
			cResponse, _ := json.Marshal(gin.H{
				"sender": sender.UUID,
				"content": content,
			})
			client.Send <- cResponse
			sender.Send <- []byte(helpers.SUCCESS)
			return
		}
	}
	// Offline
	sender.Send <- []byte(helpers.ERROR)
}