package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tLALOck64/microservicio-cuentos/internal/shared/response"
	"github.com/tLALOck64/microservicio-cuentos/internal/story/application"
)

type GetController struct {
	GetUseCase *application.GetUseCase
}

func NewGetController(getUseCase *application.GetUseCase) *GetController{
	return &GetController{GetUseCase: getUseCase}
}

func (ctr *GetController) Run(ctx *gin.Context){
	storys, err := ctr.GetUseCase.Run()

	if err != nil{
		ctx.JSON(http.StatusBadRequest, response.Response{
			Success: false,
			Message: "Error retrieved story's",
			Data: nil,
			Error: err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		Success: true,
		Message: "Successfully retrieved story's",
		Data: storys,
		Error: nil,
	})
}