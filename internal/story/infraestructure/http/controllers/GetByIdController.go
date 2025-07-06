package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tLALOck64/microservicio-cuentos/internal/shared/response"
	"github.com/tLALOck64/microservicio-cuentos/internal/story/application"
)

type GetByIdController struct {
	GetByIdUseCase *application.GetByIdUseCase
}

func NewGetByIdController(getByIdUseCase *application.GetByIdUseCase) *GetByIdController{
	return &GetByIdController{GetByIdUseCase: getByIdUseCase}
}

func (ctr *GetByIdController) Run(ctx *gin.Context){
	id := ctx.Param("id")

	if id == " "{
		ctx.JSON(http.StatusBadRequest, response.Response{
			Success: false,
			Message:"",
			Data:nil,
			Error: nil,
		})
		return
	}

	story, err := ctr.GetByIdUseCase.Run(id)

	if err != nil{
		ctx.JSON(http.StatusInternalServerError, response.Response{
			Success: false,
			Message: "error retrieved story",
			Data: nil,
			Error: err.Error(),
		})
		return 
	}

	ctx.JSON(http.StatusOK, response.Response{
		Success: true,
		Message: "Successfully retrieved story",
		Data: story,
		Error: nil,
	})
}