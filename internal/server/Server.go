package server

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tLALOck64/microservicio-cuentos/internal/story/infraestructure/http/routes"
)

type Server struct {
	engine *gin.Engine
	host string
	port string
	httpAddr string
}

func NewServer(host, port string) Server{
	gin.SetMode(gin.ReleaseMode)

	srv := Server{
		engine: gin.New(),
		host: host,
		port: port,
		httpAddr: host + ":" + port,
	}

	srv.engine.Use(func(c *gin.Context) {
		
	if c.Request.Host != srv.httpAddr {
		  c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid host header"})
		  return
		}
		c.Header("X-Frame-Options", "DENY")
		c.Header("Content-Security-Policy", "default-src 'self'; connect-src *; font-src *; script-src-elem * 'unsafe-inline'; img-src * data:; style-src * 'unsafe-inline';")
		c.Header("X-XSS-Protection", "1; mode=block")
		c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")
		c.Header("Referrer-Policy", "strict-origin")
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("Permissions-Policy", "geolocation=(),midi=(),sync-xhr=(),microphone=(),camera=(),magnetometer=(),gyroscope=(),fullscreen=(self),payment=()")
		c.Next()
	})

	srv.engine.Use(gin.Recovery())
	srv.engine.Use(gin.Logger())
	srv.engine.RedirectTrailingSlash = true

	srv.engine.GET("ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message" : "pong!",
		})
	})

	storyRoute := srv.engine.Group("v1/story")

	routes.Routes(storyRoute)


	return srv
}

func (s *Server) Run() error{
	log.Println("Starting server on" + s.httpAddr)
	return s.engine.Run(s.httpAddr)
}