package ws

import (
	"github.com/gin-gonic/gin"
	"github.com/lunny/log"
	"gin-web/app/common/middlewares"
	"github.com/gorilla/websocket"
	"time"
	"net/http"
	"github.com/golibs/uuid"
	"gin-web/app/modules/ws/controllers"
	middlewares2 "gin-web/app/modules/ws/middlewares"
)

func wsRouter(r *middlewares2.Manager) {
	// Register router
	chat := &controllers.Chat{}
	r.GET("chat", chat.Person)
}

func Routers(router *gin.Engine) {
	ws := websocket.Upgrader{
		ReadBufferSize: 1024,
		WriteBufferSize: 1024,
		HandshakeTimeout: 5 * time.Second,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	router.GET("/ws/:token", func(ctx *gin.Context) {
		token := ctx.Param("token")
		env := &middlewares.Env{}
		user, err := env.GetUserInfo(token)
		if err != nil {
			log.Println(err)
			return
		}
		conn, err := ws.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(conn)
		m := &middlewares2.Manager{
			Register: make(chan *middlewares2.Client),
			Unregister: make(chan *middlewares2.Client),
			Broadcast: make(chan []byte),
			Clients: make(map[*middlewares2.Client]bool),
			Router: make(map[string]gin.HandlerFunc),
		}
		wsRouter(m)

		go m.New()
		client := &middlewares2.Client{User: user, UUID: uuid.Rand().Hex(), Conn: conn, Send: make(chan []byte)}
		m.Register <- client

		go m.Read(client)
		go m.Write(client)
	})

}