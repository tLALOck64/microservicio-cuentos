package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/tLALOck64/microservicio-cuentos/internal/question/application"
	"github.com/tLALOck64/microservicio-cuentos/internal/question/infraestructure/http/mapper"
	"github.com/tLALOck64/microservicio-cuentos/internal/question/infraestructure/http/request"
	"github.com/tLALOck64/microservicio-cuentos/internal/shared/response"
)

type CreateController struct {
	CreateUseCase *application.CreateUseCase
	Validator     *validator.Validate
}

func NewCreateController(createUseCase *application.CreateUseCase) *CreateController {
	return &CreateController{
		CreateUseCase: createUseCase,
		Validator:     validator.New(),
	}
}

func (ctr *CreateController) Run(ctx *gin.Context) {
	var req request.CreateQuestionRequest

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response{
			Success: false,
			Message: "Error datos faltantes",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	if err := ctr.Validator.Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response{
			Success: false,
			Message: "",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	questionEntity, err := mapper.MapCreateQuestionRequest(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response{
			Success: false,
			Message: "Error de tipos",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	question, err := ctr.CreateUseCase.Run(&questionEntity)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Response{
			Success: false,
			Message: "Error al crear pregunta",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, response.Response{
		Success: true,
		Message: "Pregunta creada de manera exitosa",
		Data:    question,
		Error:   nil,
	})
}
