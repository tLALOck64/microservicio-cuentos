package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tLALOck64/microservicio-cuentos/internal/shared/response"
	"github.com/tLALOck64/microservicio-cuentos/internal/question/application"
)

type GetByStoryIdController struct {
	GetByStoryIdUseCase *application.GetByStoryIdUseCase
}

func NewGetByStoryIdController(getByStoryIdUseCase *application.GetByStoryIdUseCase) *GetByStoryIdController {
	return &GetByStoryIdController{GetByStoryIdUseCase: getByStoryIdUseCase}
}

func (ctr *GetByStoryIdController) Run(ctx *gin.Context) {
	storyID := ctx.Param("storyId")

	if storyID == "" {
		ctx.JSON(http.StatusBadRequest, response.Response{
			Success: false,
			Message: "ID de historia requerido",
			Data:    nil,
			Error:   nil,
		})
		return
	}

	questions, err := ctr.GetByStoryIdUseCase.Run(storyID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Response{
			Success: false,
			Message: "Error al obtener preguntas de la historia",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		Success: true,
		Message: "Preguntas de la historia obtenidas exitosamente",
		Data:    questions,
		Error:   nil,
	})
} 