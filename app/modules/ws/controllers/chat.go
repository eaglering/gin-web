package controllers

import (
	"github.com/gin-gonic/gin"
	"gin-web/app/modules/ws/middlewares/manager"
	"github.com/lunny/log"
)

type Chat struct{}

func (c *Chat) Person(ctx *gin.Context) {
	from := ctx.GetString("from")
	sender := ctx.GetString("sender")
	recipient := ctx.GetString("recipient")
	content := ctx.GetString("content")

	m := manager.Instance()
	for client, _ := range m.Clients {
		if client.UUID == recipient {
			// todo
			log.Println(from)
			log.Println(sender)
			log.Println(content)
			client.Send <- []byte(content)
			log.Println(client)
			break
		}
	}
}