package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tLALOck64/microservicio-cuentos/internal/shared/middleware"
	"github.com/tLALOck64/microservicio-cuentos/internal/story/infraestructure/http"
)

func Routes(router *gin.RouterGroup){
	createController := http.SetUpCreate()
	getController := http.SetUpGet()
	getByIdController := http.SetUpGetById()

	router.Use(middleware.JWTAuthMiddleware())
	router.POST("/", createController.Run)
	router.GET("/", getController.Run)
	router.GET("/:id", getByIdController.Run)
}