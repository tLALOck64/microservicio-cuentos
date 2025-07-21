package http

import (
	"github.com/tLALOck64/microservicio-cuentos/internal/question/application"
	"github.com/tLALOck64/microservicio-cuentos/internal/question/domain/ports"
	adapters "github.com/tLALOck64/microservicio-cuentos/internal/question/infraestructure/adapter"
	"github.com/tLALOck64/microservicio-cuentos/internal/question/infraestructure/http/controllers"
)

var QuestionRepository ports.QuestionRepository

func init() {
	var err error
	QuestionRepository, err = adapters.NewQuestionRepositoryMongoDB()

	if err != nil {
		panic("error al inicializar el adapter de preguntas")
	}
}

func SetUpCreate() *controllers.CreateController {
	createUseCase := application.NewCreateUseCase(QuestionRepository)
	return controllers.NewCreateController(createUseCase)
}

func SetUpGet() *controllers.GetController {
	getUseCase := application.NewGetUseCase(QuestionRepository)
	return controllers.NewGetController(getUseCase)
}

func SetUpGetById() *controllers.GetByIdController {
	getByIdUseCase := application.NewGetByIdUseCase(QuestionRepository)
	return controllers.NewGetByIdController(getByIdUseCase)
}

func SetUpGetByStoryId() *controllers.GetByStoryIdController {
	getByStoryIdUseCase := application.NewGetByStoryIdUseCase(QuestionRepository)
	return controllers.NewGetByStoryIdController(getByStoryIdUseCase)
}
