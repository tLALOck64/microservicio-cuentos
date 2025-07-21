package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tLALOck64/microservicio-cuentos/internal/shared/response"
	"github.com/tLALOck64/microservicio-cuentos/internal/story/application"
)

type GetByCategoryController struct {
	GetByCategoryUseCase *application.GetByCategoryUseCase
}

func NewGetByCategoryController(getByCategoryUseCase *application.GetByCategoryUseCase) *GetByCategoryController {
	return &GetByCategoryController{GetByCategoryUseCase: getByCategoryUseCase}
}

func (ctr *GetByCategoryController) Run(ctx *gin.Context) {
	category := ctx.Param("category")

	if category == "" {
		ctx.JSON(http.StatusBadRequest, response.Response{
			Success: false,
			Message: "Categoría requerida",
			Data:    nil,
			Error:   nil,
		})
		return
	}

	stories, err := ctr.GetByCategoryUseCase.Run(category)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Response{
			Success: false,
			Message: "Error al obtener historias por categoría",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		Success: true,
		Message: "Historias obtenidas exitosamente",
		Data:    stories,
		Error:   nil,
	})
}
