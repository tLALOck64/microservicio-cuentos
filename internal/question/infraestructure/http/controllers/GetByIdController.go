package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tLALOck64/microservicio-cuentos/internal/question/application"
	"github.com/tLALOck64/microservicio-cuentos/internal/shared/response"
)

type GetByIdController struct {
	GetByIdUseCase *application.GetByIdUseCase
}

func NewGetByIdController(getByIdUseCase *application.GetByIdUseCase) *GetByIdController {
	return &GetByIdController{GetByIdUseCase: getByIdUseCase}
}

func (ctr *GetByIdController) Run(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		ctx.JSON(http.StatusBadRequest, response.Response{
			Success: false,
			Message: "ID requerido",
			Data:    nil,
			Error:   nil,
		})
		return
	}

	question, err := ctr.GetByIdUseCase.Run(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Response{
			Success: false,
			Message: "Error al obtener pregunta",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		Success: true,
		Message: "Pregunta obtenida exitosamente",
		Data:    question,
		Error:   nil,
	})
}
