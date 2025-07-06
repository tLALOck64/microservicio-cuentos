package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/tLALOck64/microservicio-cuentos/internal/shared/response"
	"github.com/tLALOck64/microservicio-cuentos/internal/story/application"
	"github.com/tLALOck64/microservicio-cuentos/internal/story/infraestructure/http/mapper"
	"github.com/tLALOck64/microservicio-cuentos/internal/story/infraestructure/http/request"
)

type CreateController struct {
	CreateUseCase *application.CreateUseCase
	Validator *validator.Validate
}

func NewCreateUseCase(createUseCase *application.CreateUseCase) *CreateController{
	return &CreateController{
		CreateUseCase: createUseCase,
		Validator: validator.New(),
	}
}

func (ctr *CreateController) Run(ctx *gin.Context){
	var req request.CreateStoryRequest
	
	if err := ctx.ShouldBind(&req); err != nil{
		ctx.JSON(http.StatusBadRequest, response.Response{
			Success: false,
			Message: "Error datos faltantes",
			Data: nil,
			Error: err.Error(),
		})
		return
	}

	if err := ctr.Validator.Struct(req); err != nil{
		ctx.JSON(http.StatusBadRequest, response.Response{
			Success: false,
			Message: "",
			Data: nil,
			Error: err.Error(),
		})
		return
	}

	storyEntity, err := mapper.MapCreateStoryRequest(req)

	if err != nil{
		ctx.JSON(http.StatusBadRequest, response.Response{
			Success: false,
			Message: "Error de tipos",
			Data: nil,
			Error: err.Error(),
		})
		return
	}

	story, err := ctr.CreateUseCase.Run(&storyEntity)

	if err != nil{
		ctx.JSON(http.StatusInternalServerError, response.Response{
			Success: false,
			Message: "Error al crear",
			Data: nil,
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, response.Response{
		Success: true,
		Message: "Historia creada de manera exitosa",
		Data: story,
		Error: nil,
	})

	
}