package ws

import (
	"github.com/gin-gonic/gin"
	"github.com/lunny/log"
	"gin-web/app/modules/ws/middlewares/manager"
	"gin-web/app/modules/api/middlewares"
	"github.com/gorilla/websocket"
	"time"
	"net/http"
	"github.com/golibs/uuid"
	"gin-web/app/modules/ws/controllers"
)

func Routers(router *gin.Engine) {
	ws := websocket.Upgrader{
		ReadBufferSize: 1024,
		WriteBufferSize: 1024,
		HandshakeTimeout: 5 * time.Second,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	env := &middlewares.Env{}
	chat := &controllers.Chat{}
	router.GET("/ws", env.New(), func(ctx *gin.Context) {
		conn, err := ws.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			log.Println(err)
			return
		}
		in, exists := ctx.Get("Env")
		if !exists {
			log.Println("Unknown environment in request header")
			return
		}
		e := in.(middlewares.Env)
		client := &manager.Client{Uid: e.User["uid"], UUID: uuid.Rand().Hex(), Conn: conn, Send: make(chan []byte)}

		m := manager.Instance()
		// Register router
		m.GET("chat", chat.Person)
		m.Register <- client

		go client.Read()
		go client.Write()
	})

}