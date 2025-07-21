package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tLALOck64/microservicio-cuentos/internal/question/application"
	"github.com/tLALOck64/microservicio-cuentos/internal/shared/response"
)

type GetController struct {
	GetUseCase *application.GetUseCase
}

func NewGetController(getUseCase *application.GetUseCase) *GetController {
	return &GetController{GetUseCase: getUseCase}
}

func (ctr *GetController) Run(ctx *gin.Context) {
	questions, err := ctr.GetUseCase.Run()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response{
			Success: false,
			Message: "Error al obtener preguntas",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		Success: true,
		Message: "Preguntas obtenidas exitosamente",
		Data:    questions,
		Error:   nil,
	})
}
