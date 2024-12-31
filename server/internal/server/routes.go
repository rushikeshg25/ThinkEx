package server

import (
	"net/http"

	"fmt"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/coder/websocket"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Add your frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true, // Enable cookies/auth
	}))

	r.GET("/", s.HelloWorldHandler)

	r.GET("/health", s.healthHandler)

	r.GET("/websocket", s.websocketHandler)

	api:=r.Group("/api")
	{
		// api.Use()
		Auth(api.Group("/auth"))
	}

	return r
}


func Auth(g* gin.RouterGroup){
	g.GET("/login",func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "login",
		})
	})
	g.GET("/logout",)
	g.GET("/signup",)
}

func (s *Server) HelloWorldHandler(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	c.JSON(http.StatusOK, resp)
}

func (s *Server) healthHandler(c *gin.Context) {
	// Db.Create(&database.User{
	// 	Id:        "1",
	// 	Role:      "admin",
	// 	Balance:   100,
	// 	CreatedAt: time.Now(),
	// 	UpdatedAt: time.Now(),
	// })
	c.JSON(http.StatusOK, s.db.Health())
}

func (s *Server) websocketHandler(c *gin.Context) {
	w := c.Writer
	r := c.Request
	socket, err := websocket.Accept(w, r, nil)

	if err != nil {
		log.Printf("could not open websocket: %v", err)
		_, _ = w.Write([]byte("could not open websocket"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer socket.Close(websocket.StatusGoingAway, "server closing websocket")

	ctx := r.Context()
	socketCtx := socket.CloseRead(ctx)

	for {
		payload := fmt.Sprintf("server timestamp: %d", time.Now().UnixNano())
		err := socket.Write(socketCtx, websocket.MessageText, []byte(payload))
		if err != nil {
			break
		}
		time.Sleep(time.Second * 2)
	}
}
